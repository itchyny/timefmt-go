package timefmt_test

import "strings"

type stringBuilder struct {
	strings.Builder
}

func (sb *stringBuilder) writeDiff(s string) {
	sb.WriteString("\x1b[1;4m")
	sb.WriteString(s)
	sb.WriteString("\x1b[0m")
}

func diff(expected, got string) string {
	var sbx, sby stringBuilder
	xs := strings.Split(expected, " ")
	ys := strings.Split(got, " ")
	for i, j := 0, 0; ; i, j = i+1, j+1 {
		if i >= len(xs) {
			if j >= len(ys) {
				break
			}
			if j > 0 {
				sby.writeDiff(" ")
			}
			sby.writeDiff(ys[j])
			continue
		} else if j >= len(ys) {
			if i > 0 {
				sbx.writeDiff(" ")
			}
			sbx.writeDiff(xs[i])
			continue
		}
		if xs[i] == "" {
			if ys[j] != "" {
				sbx.writeDiff(" ")
				j--
				continue
			}
		} else if ys[j] == "" {
			sby.writeDiff(" ")
			i--
			continue
		}
		if i > 0 {
			sbx.WriteByte(' ')
		}
		if j > 0 {
			sby.WriteByte(' ')
		}
		if xs[i] == ys[j] {
			sbx.WriteString(xs[i])
			sby.WriteString(ys[j])
		} else {
			sbx.writeDiff(xs[i])
			sby.writeDiff(ys[j])
		}
	}
	return "diff:\nexpected: " + sbx.String() + "\n     got: " + sby.String()
}
