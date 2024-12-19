package assert

import (
	"os"
	"strconv"
	"strings"
)

func numberPadder(digits int) func(int) string {
	return func(n int) string {
		number := strconv.Itoa(n)
		return strings.Repeat("0", digits-len(number)) + number
	}
}

func get_context(f frame) string {
	start := f.Line - ContextWindow
	end := f.Line + ContextWindow
	pad := numberPadder(len(strconv.Itoa(end)))

	file, err := os.ReadFile(f.File)
	if err != nil {
		return "\tContext not available\n"
	}

	lines := strings.Split(string(file), "\n")

	var builder strings.Builder
	for i := start; i <= end; i++ {
		if i < 1 || i > len(lines) {
			continue
		}

		line := lines[i-1]
		if i == f.Line {
			builder.WriteString(
				applyAnsi(LineColor, "\t"+pad(i)+" > "+line+"\n"),
			)
		} else {
			builder.WriteString(
				applyAnsi(ContextColor, "\t"+pad(i)+" | "+line+"\n"),
			)
		}
	}

	return builder.String()
}
