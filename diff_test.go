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
	xs, ys := split(expected), split(got)
	var sbx, sby stringBuilder
	for i, j := 0, 0; i < len(xs) || j < len(ys); {
		switch {
		case j >= len(ys) || i < len(xs) && isSpaces(xs[i]) && !isSpaces(ys[j]):
			sbx.writeDiff(xs[i])
			i++
		case i >= len(xs) || isSpaces(ys[j]) && !isSpaces(xs[i]):
			sby.writeDiff(ys[j])
			j++
		case xs[i] == ys[j]:
			sbx.WriteString(xs[i])
			sby.WriteString(ys[j])
			i++
			j++
		default:
			sbx.writeDiff(xs[i])
			sby.writeDiff(ys[j])
			i++
			j++
		}
	}
	return "diff:\nexpected: " + sbx.String() + "\n     got: " + sby.String()
}

func split(s string) []string {
	var ss []string
	for i := 0; i < len(s); {
		j := i + 1
		for j < len(s) && (s[j] == ' ') == (s[i] == ' ') {
			j++
		}
		ss = append(ss, s[i:j])
		i = j
	}
	return ss
}

func isSpaces(s string) bool {
	return s[0] == ' '
}
