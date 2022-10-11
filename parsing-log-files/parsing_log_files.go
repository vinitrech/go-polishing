package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	result := regexp.MustCompile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`)
	return result.MatchString(text)
}

func SplitLogLine(text string) []string {
	expression := regexp.MustCompile(`<[\*\~\-\=]*>`)

	return expression.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	expression := regexp.MustCompile(`(?i)(\".*(password).*\")`)

	for _, line := range lines {
		if expression.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	expression := regexp.MustCompile(`end-of-line[0-9]*`)
	return expression.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	expression := regexp.MustCompile(`User\s+(\w+)`)

	for index, line := range lines {
		matches := expression.FindStringSubmatch(line)

		if len(matches) > 0 {
			lines[index] = "[USR] " + matches[1] + " " + lines[index]
		}
	}

	return lines
}
