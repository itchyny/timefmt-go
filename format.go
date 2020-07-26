package timefmt

import (
	"bytes"
	"strconv"
	"strings"
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
			if i++; i == len(format) {
				buf.WriteByte(b)
				break
			}
			b, width, padding, upper = format[i], 0, '0', false
		L:
			switch b {
			case '-':
				if pending != "" {
					buf.WriteByte(b)
					break
				}
				if i++; i == len(format) {
					appendLast(buf, format, width, padding)
					break
				}
				padding = ^paddingMask
				b = format[i]
				goto L
			case '_':
				if i++; i == len(format) {
					appendLast(buf, format, width, padding)
					break
				}
				padding = ' ' | ^paddingMask
				b = format[i]
				goto L
			case '^':
				if i++; i == len(format) {
					appendLast(buf, format, width, padding)
					break
				}
				upper = true
				b = format[i]
				goto L
			case '0':
				if i++; i == len(format) {
					appendLast(buf, format, width, padding)
					break
				}
				padding = '0' | ^paddingMask
				b = format[i]
				goto L
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				width = int(b & 0x0F)
				for i++; i < len(format); i++ {
					b = format[i]
					if b <= '9' && '0' <= b {
						width = width*10 + int(b&0x0F)
					} else {
						break
					}
				}
				if padding == ^paddingMask {
					padding = ' ' | ^paddingMask
				}
				if i == len(format) {
					appendLast(buf, format, width, padding)
					break
				}
				goto L
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
				appendString(buf, longMonthNames[month-1], width, padding, upper)
			case 'b', 'h':
				appendString(buf, shortMonthNames[month-1], width, padding, upper)
			case 'A':
				appendString(buf, longWeekNames[t.Weekday()], width, padding, upper)
			case 'a':
				appendString(buf, shortWeekNames[t.Weekday()], width, padding, upper)
			case 'w':
				for ; width > 1; width-- {
					buf.WriteByte(padding & paddingMask)
				}
				buf.WriteByte('0' + byte(t.Weekday()))
			case 'u':
				w := int(t.Weekday())
				if w == 0 {
					w = 7
				}
				for ; width > 1; width-- {
					buf.WriteByte(padding & paddingMask)
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
				if padding < ^paddingMask {
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
				if padding < ^paddingMask {
					padding = ' '
				}
				fallthrough
			case 'H':
				if width < 2 {
					width = 2
				}
				appendInt(buf, hour, width, padding)
			case 'l':
				if padding < ^paddingMask {
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
				} else if h == 0 {
					h = 12
				}
				if width < 2 {
					width = 2
				}
				appendInt(buf, h, width, padding)
			case 'p':
				if hour < 12 {
					appendString(buf, "AM", width, padding, upper)
				} else {
					appendString(buf, "PM", width, padding, upper)
				}
			case 'P':
				if hour < 12 {
					appendString(buf, "am", width, padding, upper)
				} else {
					appendString(buf, "pm", width, padding, upper)
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
				if padding < ^paddingMask {
					padding = ' '
				}
				appendInt(buf, int(t.Unix()), width, padding)
			case 'f':
				if width == 0 {
					width = 6
				}
				appendInt(buf, t.Nanosecond()/1000, width, padding)
			case 'Z', 'z':
				name, offset := t.Zone()
				if b == 'Z' && name != "" {
					appendString(buf, name, width, padding, false)
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
				appendString(buf, "\t", width, padding, false)
			case 'n':
				appendString(buf, "\n", width, padding, false)
			case '%':
				appendString(buf, "%", width, padding, false)
			default:
				if pending == "" {
					var ok bool
					if pending, ok = compositions[b]; ok {
						break
					}
					appendLast(buf, format[:i], width-1, padding)
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
	if padding != ^paddingMask {
		padding &= paddingMask
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

func appendString(buf *bytes.Buffer, str string, width int, padding byte, upper bool) {
	if width > len(str) && padding != ^paddingMask {
		if padding < ^paddingMask {
			padding = ' '
		} else {
			padding &= paddingMask
		}
		for width -= len(str); width > 0; width-- {
			buf.WriteByte(padding)
		}
	}
	if upper {
		for _, b := range []byte(str) {
			buf.WriteByte(b & 0x5F)
		}
	} else {
		buf.WriteString(str)
	}
}

func appendLast(buf *bytes.Buffer, format string, width int, padding byte) {
	appendString(buf, format[strings.LastIndexByte(format, '%'):], width, padding, false)
}

const paddingMask byte = 0x7F

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
	'T': "H:M:S",
	'X': "H:M:S",
	'r': "I:M:S p",
	'R': "H:M",
}
