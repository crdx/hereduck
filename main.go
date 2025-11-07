package hereduck

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

// D strips leading indentation from a string. The minimum indentation level shared by all lines
// is removed, but any subsequent indentation is preserved.
func D(str string) string {
	if len(str) > 0 && str[0] == '\n' {
		str = str[1:]
	}

	lines := strings.Split(str, "\n")
	min := getMinIndentationLevel(lines)

	// If the length of the last line is smaller or equal to the minimum indentation level then
	// remove the indentation completely. This preserves the trailing newline without the excess
	// whitespace.
	if len(lines[len(lines)-1]) <= min {
		lines[len(lines)-1] = ""
	}

	lines = removeIndentation(lines, min)
	return strings.Join(lines, "\n")
}

// Df is like [D] but also follows the standard formatting string pattern.
func Df(format string, a ...any) string {
	return fmt.Sprintf(D(format), a...)
}

// —————————————————————————————————————————————————————————————————————————————————————————————————

func getMinIndentationLevel(lines []string) int {
	min := math.MaxInt32

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		current := 0

		// Calculate the number of whitespace characters leading this line.
		for _, rune := range line {
			if unicode.IsSpace(rune) {
				current++
			} else {
				break
			}
		}

		// If this is the last line and it's only whitespace, skip it. No point in counting it as it
		// will be removed later on.
		if len(line) == current && i == len(lines)-1 {
			continue
		}

		if current < min {
			min = current
		}
	}

	return min
}

func removeIndentation(lines []string, n int) []string {
	for i, line := range lines {
		if len(lines[i]) >= n {
			lines[i] = line[n:]
		}
	}

	return lines
}
