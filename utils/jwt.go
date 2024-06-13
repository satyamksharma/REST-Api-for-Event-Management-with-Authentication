package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)
const secretKey = "JaiShreeRam"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	})

	return token.SignedString([]byte(secretKey))

}
