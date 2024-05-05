package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func (Validator) ValidateMD5(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile("^[0-9a-fA-F]{32}$")
	return re.MatchString(value)
}
