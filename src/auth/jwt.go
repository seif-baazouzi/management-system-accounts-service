package auth

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(key string, value string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[key] = value

	secret := []byte(os.Getenv("JWT_SECRET"))
	t, _ := token.SignedString(secret)
	return t
}
