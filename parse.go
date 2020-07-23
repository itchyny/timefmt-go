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
	year, month, day, hour, min, sec, nsec, loc := 0, 1, 1, 0, 0, 0, 0, time.UTC
	defer func() {
		if err != nil {
			err = &parseError{source, format, err}
		}
	}()
	var j int
	for i, l := 0, len(source); i < len(format); i++ {
		if b := format[i]; b == '%' {
			i++
			b = format[i]
			switch b {
			case 'Y':
				if j+4 > l {
					err = errors.New("cannot parse %Y")
					return
				}
				if year, err = strconv.Atoi(string(source[j : j+4])); err != nil {
					return
				}
				j += 4
			case 'm':
				if j >= l {
					err = errors.New("cannot parse %m")
					return
				}
				if c := source[j]; c > '1' || j+1 == l || !isDigit(source[j+1]) {
					if month, err = strconv.Atoi(string(source[j : j+1])); err != nil {
						return
					}
					j++
				} else if j+2 <= l {
					if month, err = strconv.Atoi(string(source[j : j+2])); err != nil {
						return
					}
					j += 2
				} else {
					err = errors.New("cannot parse %m")
					return
				}
			case 'd':
				if j >= l {
					err = errors.New("cannot parse %d")
					return
				}
				if c := source[j]; c > '1' || j+1 == l || !isDigit(source[j+1]) {
					if day, err = strconv.Atoi(string(source[j : j+1])); err != nil {
						return
					}
					j++
				} else if j+2 <= l {
					if day, err = strconv.Atoi(string(source[j : j+2])); err != nil {
						return
					}
					j += 2
				} else {
					err = errors.New("cannot parse %d")
					return
				}
			default:
				err = fmt.Errorf("unexpected format: %q", format[i-1:i+1])
				return
			}
		} else if j >= len(source) || b != source[j] {
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
	return time.Date(year, time.Month(month), day, hour, min, sec, nsec, loc), nil
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
