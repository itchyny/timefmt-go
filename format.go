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
	var width int
	var padding byte
	var pending string
	var upper bool
	for i := 0; i < len(format); i++ {
		if b := format[i]; b == '%' {
			i++
			if i == len(format) {
				buf.WriteByte(b)
				break
			}
			b = format[i]
			width, padding, upper = 0, '0', false
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
			} else if b == '^' {
				upper = true
				i++
				if i == len(format) {
					buf.WriteByte(b)
					break
				}
				b = format[i]
			}
			if b <= '9' && '1' <= b {
				width = int(b & 0x0F)
				for i++; i < len(format); i++ {
					b = format[i]
					if b <= '9' && '0' <= b {
						width = width*10 + int(b&0x0F)
					} else {
						break
					}
				}
				if i == len(format) {
					appendString(buf, "%"+strconv.Itoa(width), width, false)
					break
				}
				if padding == 0 {
					padding = ' '
				}
			}
		L:
			switch b {
			case 'Y':
				if width == 0 {
					width = 4
				}
				appendInt(buf, year, width, padding)
			case 'y':
				if width < 2 {
					width = 2
				}
				appendInt(buf, year%100, width, padding)
			case 'C':
				if width < 2 {
					width = 2
				}
				appendInt(buf, year/100, width, padding)
			case 'g':
				year, _ := t.ISOWeek()
				if width < 2 {
					width = 2
				}
				appendInt(buf, year%100, width, padding)
			case 'G':
				year, _ := t.ISOWeek()
				if width == 0 {
					width = 4
				}
				appendInt(buf, year, width, padding)
			case 'm':
				if width < 2 {
					width = 2
				}
				appendInt(buf, int(month), width, padding)
			case 'B':
				appendString(buf, longMonthNames[month-1], width, upper)
			case 'b', 'h':
				appendString(buf, shortMonthNames[month-1], width, upper)
			case 'A':
				appendString(buf, longWeekNames[t.Weekday()], width, upper)
			case 'a':
				appendString(buf, shortWeekNames[t.Weekday()], width, upper)
			case 'w':
				for ; width > 1; width-- {
					buf.WriteByte(padding)
				}
				buf.WriteByte('0' + byte(t.Weekday()))
			case 'u':
				w := int(t.Weekday())
				if w == 0 {
					w = 7
				}
				for ; width > 1; width-- {
					buf.WriteByte(padding)
				}
				buf.WriteByte('0' + byte(w))
			case 'V':
				_, week := t.ISOWeek()
				if width < 2 {
					width = 2
				}
				appendInt(buf, week, width, padding)
			case 'U':
				week := (t.YearDay() + 6 - int(t.Weekday())) / 7
				if width < 2 {
					width = 2
				}
				appendInt(buf, week, width, padding)
			case 'W':
				week := t.YearDay()
				if int(t.Weekday()) > 0 {
					week -= int(t.Weekday()) - 7
				}
				week /= 7
				if width < 2 {
					width = 2
				}
				appendInt(buf, week, width, padding)
			case 'e':
				if padding != 0 {
					padding = ' '
				}
				fallthrough
			case 'd':
				if width < 2 {
					width = 2
				}
				appendInt(buf, day, width, padding)
			case 'j':
				if width < 3 {
					width = 3
				}
				appendInt(buf, t.YearDay(), width, padding)
			case 'k':
				if padding != 0 {
					padding = ' '
				}
				fallthrough
			case 'H':
				if width < 2 {
					width = 2
				}
				appendInt(buf, hour, width, padding)
			case 'l':
				if padding != 0 {
					padding = ' '
				}
				h := hour
				if h > 12 {
					h -= 12
				}
				if width < 2 {
					width = 2
				}
				appendInt(buf, h, width, padding)
			case 'I':
				h := hour
				if h > 12 {
					h -= 12
				}
				if h == 0 {
					h = 12
				}
				if width < 2 {
					width = 2
				}
				appendInt(buf, h, width, padding)
			case 'p':
				for ; width > 2; width-- {
					buf.WriteByte(' ')
				}
				if hour <= 12 {
					buf.WriteString("AM")
				} else {
					buf.WriteString("PM")
				}
			case 'P':
				for ; width > 2; width-- {
					buf.WriteByte(' ')
				}
				if hour <= 12 {
					buf.WriteString("am")
				} else {
					buf.WriteString("pm")
				}
			case 'M':
				if width < 2 {
					width = 2
				}
				appendInt(buf, min, width, padding)
			case 'S':
				if width < 2 {
					width = 2
				}
				appendInt(buf, sec, width, padding)
			case 's':
				appendInt(buf, int(t.Unix()), width, ' ')
			case 'f':
				if width == 0 {
					width = 6
				}
				appendInt(buf, t.Nanosecond()/1000, width, padding)
			case 'Z', 'z':
				name, offset := t.Zone()
				if b == 'Z' && name != "" {
					appendString(buf, name, width, false)
					break
				}
				if offset < 0 {
					buf.WriteRune('-')
					offset = -offset
				} else {
					buf.WriteRune('+')
				}
				offset /= 60
				if width < 4 {
					width = 4
				}
				appendInt(buf, (offset/60)*100+offset%60, width, padding)
			case 't':
				appendString(buf, "\t", width, false)
			case 'n':
				appendString(buf, "\n", width, false)
			default:
				if pending == "" {
					var ok bool
					if pending, ok = compositions[b]; ok {
						break
					}
				}
				buf.WriteByte(b)
			}
			if pending != "" {
				b, pending, width, padding = pending[0], pending[1:], 0, '0'
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

func appendString(buf *bytes.Buffer, str string, width int, upper bool) {
	for width -= len(str); width > 0; width-- {
		buf.WriteByte(' ')
	}
	if upper {
		for _, b := range []byte(str) {
			buf.WriteByte(b & 0x5F)
		}
	} else {
		buf.WriteString(str)
	}
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
