package libraries

import "golang.org/x/crypto/bcrypt"

func ComparePassword(dbpassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(password))
}

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
