package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PasswordValidator(ft validator.FieldLevel) bool {
	password := ft.Field().String()
	var (
		hasMinLen  = len(password) >= 8
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecial = regexp.MustCompile(`[!@#~$%^&*()+|_.,<>?/{}\\-]`).MatchString(password)
	)
	return hasMinLen && hasLower && hasNumber && hasSpecial && hasUpper
}
