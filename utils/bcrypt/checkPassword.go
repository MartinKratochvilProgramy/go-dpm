package bcrypt

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
