package utils

import "strings"

func ValidateUserInputs(
	firstName string,
	lastName string,
	email string,
	userTickets int,
) (bool, bool, bool) {
	var isValidUsername bool = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets >= 0

	return isValidUsername, isValidEmail, isValidUserTickets
}
