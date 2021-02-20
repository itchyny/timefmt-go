package timefmt

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type parseError struct {
	source, format string
	err            error
}

func (err *parseError) Error() string {
	return fmt.Sprintf("failed to parse %q with %q: %s", err.source, err.format, err.err)
}

// Parse time string using the format.
func Parse(source, format string) (t time.Time, err error) {
	year, month, day, hour, min, sec, nsec, loc := 1900, 1, 1, 0, 0, 0, 0, time.UTC
	defer func() {
		if err != nil {
			err = &parseError{source, format, err}
		}
	}()
	var j, diff, century, yday int
	var pm bool
	var pending string
	for i, l := 0, len(source); i < len(format); i++ {
		if b := format[i]; b == '%' {
			i++
			if i == len(format) {
				err = errors.New("stray %")
				return
			}
			b = format[i]
		L:
			switch b {
			case 'Y':
				if year, j, err = parseNumber(source, j, 4, 'Y'); err != nil {
					return
				}
			case 'y':
				if year, j, err = parseNumber(source, j, 2, 'y'); err != nil {
					return
				}
				if year < 69 {
					year += 2000
				} else {
					year += 1900
				}
			case 'C':
				if century, j, err = parseNumber(source, j, 2, 'C'); err != nil {
					return
				}
			case 'g':
				if year, j, err = parseNumber(source, j, 2, b); err != nil {
					return
				}
				year += 2000
			case 'G':
				if year, j, err = parseNumber(source, j, 4, b); err != nil {
					return
				}
			case 'm':
				if month, j, err = parseNumber(source, j, 2, 'm'); err != nil {
					return
				}
			case 'B':
				if month, diff, err = lookup(source[j:], longMonthNames, 'B'); err != nil {
					return
				}
				j += diff
			case 'b', 'h':
				if month, diff, err = lookup(source[j:], shortMonthNames, b); err != nil {
					return
				}
				j += diff
			case 'A':
				if _, diff, err = lookup(source[j:], longWeekNames, 'A'); err != nil {
					return
				}
				j += diff
			case 'a':
				if _, diff, err = lookup(source[j:], shortWeekNames, 'a'); err != nil {
					return
				}
				j += diff
			case 'w':
				if j >= l || source[j] < '0' || '6' < source[j] {
					err = parseFormatError(b)
					return
				}
				j++
			case 'u':
				if j >= l || source[j] < '1' || '7' < source[j] {
					err = parseFormatError(b)
					return
				}
				j++
			case 'V', 'U', 'W':
				if _, j, err = parseNumber(source, j, 2, b); err != nil {
					return
				}
			case 'e':
				if j < l && source[j] == ' ' {
					j++
				}
				fallthrough
			case 'd':
				if day, j, err = parseNumber(source, j, 2, b); err != nil {
					return
				}
			case 'j':
				if yday, j, err = parseNumber(source, j, 3, 'j'); err != nil {
					return
				}
			case 'k':
				if j < l && source[j] == ' ' {
					j++
				}
				fallthrough
			case 'H':
				if hour, j, err = parseNumber(source, j, 2, b); err != nil {
					return
				}
			case 'l':
				if j < l && source[j] == ' ' {
					j++
				}
				fallthrough
			case 'I':
				if hour, j, err = parseNumber(source, j, 2, b); err != nil {
					return
				}
				if hour == 12 {
					hour = 0
				}
			case 'p', 'P':
				var ampm int
				if ampm, diff, err = lookup(source[j:], []string{"AM", "PM"}, 'p'); err != nil {
					return
				}
				j += diff
				pm = ampm == 2
			case 'M':
				if min, j, err = parseNumber(source, j, 2, 'M'); err != nil {
					return
				}
			case 'S':
				if sec, j, err = parseNumber(source, j, 2, 'S'); err != nil {
					return
				}
			case 's':
				var unix int
				if unix, j, err = parseNumber(source, j, 10, 's'); err != nil {
					return
				}
				t = time.Unix(int64(unix), 0).In(time.UTC)
				var mon time.Month
				year, mon, day = t.Date()
				hour, min, sec = t.Clock()
				month = int(mon)
			case 'f':
				var msec, k, d int
				if msec, k, err = parseNumber(source, j, 6, 'f'); err != nil {
					return
				}
				nsec = msec * 1000
				for j, d = k, k-j; d < 6; d++ {
					nsec *= 10
				}
			case 'Z':
				k := j
				for ; k < l; k++ {
					if c := source[k]; c < 'A' || 'Z' < c {
						break
					}
				}
				t, err = time.Parse("MST", source[j:k])
				if err != nil {
					err = fmt.Errorf(`cannot parse %q with "%%Z"`, source[j:k])
					return
				}
				loc = t.Location()
				j = k
			case 'z':
				if j+5 > l {
					err = parseFormatError(b)
					return
				}
				var offset int
				switch source[j] {
				case '+', '-':
					var hour, min int
					if hour, err = strconv.Atoi(source[j+1 : j+3]); err != nil {
						return
					}
					if min, err = strconv.Atoi(source[j+3 : j+5]); err != nil {
						return
					}
					offset = (hour*60 + min) * 60
					if source[j] == '-' {
						offset = -offset
					}
				default:
					err = parseFormatError(b)
					return
				}
				loc = time.FixedZone("", offset)
				j += 5
			case ':':
				if pending != "" {
					if j >= l || source[j] != b {
						err = fmt.Errorf("expected %q", b)
						return
					}
					j++
				} else {
					var hms bool
					if i++; i == len(format) {
						err = errors.New(`expected 'z' after "%:"`)
						return
					} else if b = format[i]; b == 'z' {
						if j+6 > l {
							err = errors.New("cannot parse %:z")
							return
						}
					} else {
						if b != ':' {
							err = errors.New(`expected 'z' after "%:"`)
							return
						}
						if i++; i == len(format) || format[i] != 'z' {
							err = errors.New(`expected 'z' after "%::"`)
							return
						}
						b, hms = 'z', true
						if j+9 > l {
							err = errors.New("cannot parse %::z")
							return
						}
					}
					offset := 1
					switch source[j] {
					case '-':
						offset = -1
						fallthrough
					case '+':
						var hour, min, sec int
						if hour, err = strconv.Atoi(source[j+1 : j+3]); err != nil {
							return
						}
						if source[j+3] != ':' {
							err = errors.New("expected ':' for %:z")
							return
						}
						if min, err = strconv.Atoi(source[j+4 : j+6]); err != nil {
							return
						}
						if hms {
							if source[j+6] != ':' {
								err = errors.New("expected ':' for %::z")
								return
							}
							if sec, err = strconv.Atoi(source[j+7 : j+9]); err != nil {
								return
							}
							j += 9
						} else {
							j += 6
						}
						offset *= (hour*60+min)*60 + sec
					default:
						err = parseFormatError(b)
						return
					}
					loc = time.FixedZone("", offset)
				}
			case 't', 'n':
				k := j
			K:
				for ; k < l; k++ {
					switch source[k] {
					case ' ', '\t', '\n', '\v', '\f', '\r':
					default:
						break K
					}
				}
				if k == j {
					err = fmt.Errorf("expected a space for %%%c", b)
					return
				}
				j = k
			case '%':
				if j >= l || source[j] != b {
					err = fmt.Errorf("expected %q", b)
					return
				}
				j++
			default:
				if pending == "" {
					var ok bool
					if pending, ok = compositions[b]; ok {
						break
					}
					err = fmt.Errorf(`unexpected format: "%%%c"`, b)
					return
				}
				if j >= l || source[j] != b {
					err = fmt.Errorf("expected %q", b)
					return
				}
				j++
			}
			if pending != "" {
				b, pending = pending[0], pending[1:]
				goto L
			}
		} else if j >= len(source) || source[j] != b {
			err = fmt.Errorf("expected %q", b)
			return
		} else {
			j++
		}
	}
	if j < len(source) {
		err = fmt.Errorf("unconverted string: %q", source[j:])
		return
	}
	if pm {
		hour += 12
	}
	if century > 0 {
		year = century*100 + year%100
	}
	if yday > 0 {
		return time.Date(year, time.January, 1, hour, min, sec, nsec, loc).AddDate(0, 0, yday-1), nil
	}
	return time.Date(year, time.Month(month), day, hour, min, sec, nsec, loc), nil
}

type parseFormatError byte

func (err parseFormatError) Error() string {
	return fmt.Sprintf("cannot parse %%%c", byte(err))
}

func parseNumber(source string, min, size int, format byte) (int, int, error) {
	var val int
	if l := len(source); min+size > l {
		size = l
	} else {
		size += min
	}
	i := min
	for ; i < size; i++ {
		if b := source[i]; '0' <= b && b <= '9' {
			val = val*10 + int(b&0x0F)
		} else {
			break
		}
	}
	if i == min {
		return 0, 0, parseFormatError(format)
	}
	return val, i, nil
}

func lookup(source string, candidates []string, format byte) (int, int, error) {
L:
	for i, xs := range candidates {
		for j, x := range []byte(xs) {
			if j >= len(source) {
				continue L
			}
			if y := source[j]; x != y && x|('a'-'A') != y|('a'-'A') {
				continue L
			}
		}
		return i + 1, len(xs), nil
	}
	return 0, 0, parseFormatError(format)
}
