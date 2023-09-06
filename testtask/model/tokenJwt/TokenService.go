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

func(ts *TokenService) GenerateTokens(guid string) (string, string, error) {
	accessTokenExp := time.Now().Add(time.Hour).Unix() // Время действия токена 1 час
	refreshTokenExp := time.Now().Add(72*time.Hour).Unix() // Время действия токена 3 дня

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"guid": guid,
		"exp": accessTokenExp,
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"guid": guid,
		"exp": refreshTokenExp,
	})

	accessTokenString, _ := accessToken.SignedString(ts.secretKey)
	refreshTokenString, err := refreshToken.SignedString(ts.secretKey)

	return accessTokenString, refreshTokenString, err
}

func(ts *TokenService) Valid(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ts.secretKey, nil
	})
	
	return token.Valid
}

func(ts *TokenService) GetClaims(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ts.secretKey, nil
	})
	if err != nil {
        return nil
    }

    claims, _ := token.Claims.(jwt.MapClaims)
	
	return claims
}