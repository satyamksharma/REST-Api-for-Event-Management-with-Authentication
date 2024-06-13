package utils

import (
	"errors"
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

func VerifyToken(tokenString string)  error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("Could not parse token!")
	}
	
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return errors.New("Token is not valid!")
	}

	


	return nil
}