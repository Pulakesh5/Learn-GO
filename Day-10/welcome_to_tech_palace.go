package techpalace

import (
	"fmt"
	"strings"
)

const WelcomeTemplate string = "Welcome to the Tech Palace, "

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprint(WelcomeTemplate, strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprint(stars, "\n", welcomeMsg, "\n", stars)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return fmt.Sprint(strings.TrimFunc(oldMsg, func(r rune) bool {
		return (r == rune('\n') || r == rune('*') || r == rune(' '))
	}))
}
