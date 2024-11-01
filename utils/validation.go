package utils

import "regexp"

func ValidateEmail(email string) bool {
	// Regular expression pattern for a basic email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
