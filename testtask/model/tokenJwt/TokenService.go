package tokenJwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	accessToken string
	refreshToken string
	secretKey []byte
}

func(ts *TokenService) GenerateSecretKey() { // TODO переделать тело ф-ии
	secretKey := []byte("my_secret_key")
	copy(ts.secretKey, secretKey)
}

func(ts *TokenService) GetTokens() (string, string) {
	return ts.accessToken, ts.refreshToken
}

func(ts *TokenService) GenerateToken(guid string) (string, error) {
	exp := time.Now().Add(time.Hour).Unix() // Время действия токена 1 час

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"guid": guid,
		"exp": exp,
	})

	tokenString, err := token.SignedString(ts.secretKey)

	return tokenString, err
}

func(ts *TokenService) Valid(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ts.secretKey, nil
	})
	
	return token.Valid
}