package validate

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

func (Validator) ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 || len(password) > 18 {
		return false
	}

	var hasLetter, hasNumber, hasSymbol bool
	for _, c := range password {
		switch {
		case unicode.IsLetter(c):
			hasLetter = true
		case unicode.IsDigit(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSymbol = true
		}
	}

	// 计算包含的字符类型数
	typesCount := 0
	if hasLetter {
		typesCount++
	}
	if hasNumber {
		typesCount++
	}
	if hasSymbol {
		typesCount++
	}

	return typesCount >= 2
}
