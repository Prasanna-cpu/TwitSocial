package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id string, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(60 * time.Minute).Unix(),
	})
	key := []byte(secretKey)
	tokenStr, err := token.SignedString(key)
	return tokenStr, err
}

func ValidateToken(tokenStr, secretKey string, withClaimsValidation bool) (string, string, error) {
	var (
		key    = []byte(secretKey)
		claims = jwt.MapClaims{}
		token  *jwt.Token
		err    error
	)

	if withClaimsValidation {
		token, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	} else {
		token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		}, jwt.WithoutClaimsValidation())
	}
	if err != nil {
		return "", "", err
	}
	if !token.Valid {
		return "", "", errors.New("invalid token")
	}
	return claims["id"].(string), claims["username"].(string), nil
}
