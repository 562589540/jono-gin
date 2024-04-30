package gutils

import "golang.org/x/crypto/bcrypt"

// Encrypt 密码加密
func Encrypt(stText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(stText), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CompareHashAndPassword 密码验证
func CompareHashAndPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
