package validate

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func (Validator) ValidateRuleApi(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return strings.HasPrefix(value, "/api")
}
