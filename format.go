package timefmt

import (
	"bytes"
	"strconv"
	"time"
)

// Format time to string using the format.
func Format(t time.Time, format string) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	buf := new(bytes.Buffer)
	var padding byte
	var pending string
	for i := 0; i < len(format); i++ {
		if b := format[i]; b == '%' {
			i++
			if i == len(format) {
				buf.WriteByte(b)
				break
			}
			b = format[i]
			padding = '0'
			if b == '-' || b == '_' {
				i++
				if i == len(format) {
					buf.WriteByte(b)
					break
				}
				padding = 0
				if b == '_' {
					padding = ' '
				}
				b = format[i]
			}
		L:
			switch b {
			case 'Y':
				appendInt(buf, year, 4, '0')
			case 'y':
				appendInt(buf, year%100, 2, '0')
			case 'C':
				appendInt(buf, year/100, 2, '0')
			case 'g', 'G':
				year, _ := t.ISOWeek()
				if b == 'g' {
					appendInt(buf, year%100, 2, '0')
				} else {
					appendInt(buf, year, 4, '0')
				}
			case 'm':
				appendInt(buf, int(month), 2, padding)
			case 'B':
				buf.WriteString(longMonthNames[month-1])
			case 'b', 'h':
				buf.WriteString(shortMonthNames[month-1])
			case 'A':
				buf.WriteString(longWeekNames[t.Weekday()])
			case 'a':
				buf.WriteString(shortWeekNames[t.Weekday()])
			case 'w':
				buf.WriteByte('0' + byte(t.Weekday()))
			case 'u':
				w := int(t.Weekday())
				if w == 0 {
					w = 7
				}
				buf.WriteByte('0' + byte(w))
			case 'V':
				_, week := t.ISOWeek()
				appendInt(buf, week, 2, padding)
			case 'U':
				week := (t.YearDay() + 6 - int(t.Weekday())) / 7
				appendInt(buf, week, 2, padding)
			case 'W':
				week := t.YearDay()
				if int(t.Weekday()) > 0 {
					week -= int(t.Weekday()) - 7
				}
				week /= 7
				appendInt(buf, week, 2, padding)
			case 'e':
				if padding != 0 {
					padding = ' '
				}
				fallthrough
			case 'd':
				appendInt(buf, day, 2, padding)
			case 'j':
				appendInt(buf, t.YearDay(), 3, padding)
			case 'k':
				if padding != 0 {
					padding = ' '
				}
				fallthrough
			case 'H':
				appendInt(buf, hour, 2, padding)
			case 'l':
				if padding != 0 {
					padding = ' '
				}
				h := hour
				if h > 12 {
					h -= 12
				}
				appendInt(buf, h, 2, padding)
			case 'I':
				h := hour
				if h > 12 {
					h -= 12
				}
				if h == 0 {
					h = 12
				}
				appendInt(buf, h, 2, padding)
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
				appendInt(buf, min, 2, padding)
			case 'S':
				appendInt(buf, sec, 2, padding)
			case 's':
				buf.WriteString(strconv.FormatInt(t.Unix(), 10))
			case 'f':
				appendInt(buf, t.Nanosecond()/1000, 6, '0')
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
				appendInt(buf, offset/60, 2, '0')
				appendInt(buf, offset%60, 2, '0')
			case 't':
				buf.WriteRune('\t')
			case 'n':
				buf.WriteRune('\n')
			default:
				if pending == "" {
					var ok bool
					if pending, ok = compositions[b]; ok {
						padding = '0'
						break
					}
				}
				buf.WriteByte(b)
			}
			if pending != "" {
				b, pending, padding = pending[0], pending[1:], '0'
				goto L
			}
		} else {
			buf.WriteByte(b)
		}
	}
	return buf.String()
}

func appendInt(buf *bytes.Buffer, num, width int, padding byte) {
	if padding != 0 {
		switch width {
		case 2:
			if num < 10 {
				buf.WriteByte(padding)
			}
		case 4:
			if num < 1000 {
				buf.WriteByte(padding)
				if num < 100 {
					buf.WriteByte(padding)
					if num < 10 {
						buf.WriteByte(padding)
					}
				}
			}
		default:
			str := strconv.Itoa(num)
			for width -= len(str); width > 0; width-- {
				buf.WriteByte(padding)
			}
			buf.WriteString(str)
			return
		}
	}
	buf.WriteString(strconv.Itoa(num))
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

var compositions = map[byte]string{
	'c': "a b e H:M:S Y",
	'+': "a b e H:M:S Z Y",
	'F': "Y-m-d",
	'D': "m/d/y",
	'x': "m/d/y",
	'v': "e-b-Y",
	'T': "H:M:S",
	'X': "H:M:S",
	'r': "I:M:S p",
	'R': "H:M",
}
