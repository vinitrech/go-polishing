package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var border string

	for x := 0; x < numStarsPerLine; x++ {
		border += "*"
	}

	return fmt.Sprintf("%s\n", border) + welcomeMsg + fmt.Sprintf("\n%s", border)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.Replace(oldMsg, "*", "", -1))
}
