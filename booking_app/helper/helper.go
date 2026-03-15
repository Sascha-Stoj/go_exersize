package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingConferenceTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := len(email) > 5 && strings.Contains(email, "@")
	isValidNumberUserTickets := userTickets > 0 && userTickets <= remainingConferenceTickets
	return isValidName, isValidEmail, isValidNumberUserTickets
}
