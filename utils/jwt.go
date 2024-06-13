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

func VerifyToken(tokenString string)  (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token!")
	}
	
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Token is not valid!")
	}

	
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims!")
	}

	// email := claims["email"].(string)
	userID := int64(claims["userId"].(float64))


	return userID, nil
}