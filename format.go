package timefmt

import (
	"bytes"
	"fmt"
	"time"
)

type formatError struct {
	t      time.Time
	format string
	err    error
}

func (err *formatError) Error() string {
	return fmt.Sprintf("failed to format %q with %q: %s", err.t, err.format, err.err)
}

// Format time to string using the format.
func Format(t time.Time, format string) (s string, err error) {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	defer func() {
		if err != nil {
			err = &formatError{t, format, err}
		}
	}()
	var buf bytes.Buffer
	var padZero bool
	var pending string
	for i := 0; i < len(format); i++ {
		if b := format[i]; b == '%' {
			i++
			if i == len(format) {
				buf.WriteByte(b)
				break
			}
			b = format[i]
			padZero = true
			if b == '-' {
				i++
				if i == len(format) {
					buf.WriteByte(b)
					break
				}
				padZero = false
				b = format[i]
			}
		L:
			switch b {
			case 'Y':
				if year < 1000 && padZero {
					buf.WriteRune('0')
					if year < 100 {
						buf.WriteRune('0')
						if year < 10 {
							buf.WriteRune('0')
						}
					}
				}
				buf.WriteString(fmt.Sprint(year))
			case 'y':
				y := year % 100
				if y < 10 && padZero {
					buf.WriteRune('0')
				}
				buf.WriteString(fmt.Sprint(y))
			case 'C':
				c := year / 100
				if c < 10 && padZero {
					buf.WriteRune('0')
				}
				buf.WriteString(fmt.Sprint(c))
			case 'm':
				if month < 10 && padZero {
					buf.WriteRune('0')
				}
				buf.WriteString(fmt.Sprint(int(month)))
			case 'B':
				buf.WriteString(longMonthNames[month-1])
			case 'b', 'h':
				buf.WriteString(shortMonthNames[month-1])
			case 'A':
				buf.WriteString(longWeekNames[t.Weekday()])
			case 'a':
				buf.WriteString(shortWeekNames[t.Weekday()])
			case 'w':
				buf.WriteString(fmt.Sprint(int(t.Weekday())))
			case 'u':
				w := int(t.Weekday())
				if w == 0 {
					w = 7
				}
				buf.WriteString(fmt.Sprint(w))
			case 'e':
				if day < 10 && padZero {
					buf.WriteRune(' ')
				}
				buf.WriteString(fmt.Sprint(day))
			case 'd':
				if day < 10 && padZero {
					buf.WriteRune('0')
				}
				buf.WriteString(fmt.Sprint(day))
			case 'j':
				yday := t.YearDay()
				if yday < 100 && padZero {
					buf.WriteRune('0')
					if yday < 10 {
						buf.WriteRune('0')
					}
				}
				buf.WriteString(fmt.Sprint(yday))
			case 'D', 'x':
				padZero = true
				pending = "m/d/y"
			case 'F':
				padZero = true
				pending = "Y-m-d"
			case 'v':
				padZero = true
				pending = "e-b-Y"
			case 'k':
				if hour < 10 && padZero {
					buf.WriteRune(' ')
				}
				buf.WriteString(fmt.Sprint(hour))
			case 'H':
				if hour < 10 && padZero {
					buf.WriteRune('0')
				}
				buf.WriteString(fmt.Sprint(hour))
			case 'l':
				h := hour
				if h > 12 {
					h -= 12
				}
				if h < 10 && padZero {
					buf.WriteRune(' ')
				}
				buf.WriteString(fmt.Sprint(h))
			case 'I':
				h := hour
				if h > 12 {
					h -= 12
				}
				if h < 10 && padZero {
					buf.WriteRune('0')
				}
				buf.WriteString(fmt.Sprint(h))
			case 'p':
				if hour <= 12 {
					buf.WriteString("AM")
				} else {
					buf.WriteString("PM")
				}
			case 'P':
				if hour <= 12 {
					buf.WriteString("am")
				} else {
					buf.WriteString("pm")
				}
			case 'M':
				if min < 10 && padZero {
					buf.WriteRune('0')
				}
				buf.WriteString(fmt.Sprint(min))
			case 'S':
				if sec < 10 && padZero {
					buf.WriteRune('0')
				}
				buf.WriteString(fmt.Sprint(sec))
			case 'R':
				padZero = true
				pending = "H:M"
			case 'r':
				padZero = true
				pending = "I:M:S p"
			case 'T', 'X':
				padZero = true
				pending = "H:M:S"
			case 'c':
				padZero = true
				pending = "a b e H:M:S Y"
			case 'f':
				buf.WriteString(fmt.Sprintf("%06d", t.Nanosecond()/1000))
			case 'Z', 'z':
				name, offset := t.Zone()
				if b == 'Z' && name != "" {
					buf.WriteString(name)
					break
				}
				if offset < 0 {
					buf.WriteRune('-')
					offset = -offset
				} else {
					buf.WriteRune('+')
				}
				offset /= 60
				buf.WriteString(fmt.Sprintf("%02d%02d", offset/60, offset%60))
			case 't':
				buf.WriteRune('\t')
			case 'n':
				buf.WriteRune('\n')
			default:
				buf.WriteByte(b)
			}
			if pending != "" {
				b, pending = pending[0], pending[1:]
				goto L
			}
		} else {
			buf.WriteByte(b)
		}
	}
	return buf.String(), nil
}

var longMonthNames = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

var shortMonthNames = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var longWeekNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var shortWeekNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}
