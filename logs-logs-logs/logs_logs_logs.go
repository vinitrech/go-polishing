package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, value := range log {
		if value == 10071 {
			return "recommendation"
		} else if value == 128269 {
			return "search"
		} else if value == 9728 {
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newString := ""

	for _, value := range log {
		if value == oldRune {
			newString += string(newRune)
		} else {
			newString += string(value)
		}
	}

	return newString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
