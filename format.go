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
	defer func() {
		if err != nil {
			err = &formatError{t, format, err}
		}
	}()
	var buf bytes.Buffer
	for i := 0; i < len(format); i++ {
		if b := format[i]; b == '%' {
			i++
			b = format[i]
			switch b {
			case 'Y':
				if year < 1000 {
					buf.WriteRune('0')
					if year < 100 {
						buf.WriteRune('0')
						if year < 10 {
							buf.WriteRune('0')
						}
					}
				}
				buf.WriteString(fmt.Sprint(year))
			case 'm':
				if month < 10 {
					buf.WriteRune('0')
				}
				buf.WriteString(fmt.Sprint(int(month)))
			case 'd':
				if day < 10 {
					buf.WriteRune('0')
				}
				buf.WriteString(fmt.Sprint(day))
			default:
				return "", fmt.Errorf("unexpected format: %q", format[i-1:i+1])
			}
		} else {
			buf.WriteByte(b)
		}
	}
	return buf.String(), nil
}
