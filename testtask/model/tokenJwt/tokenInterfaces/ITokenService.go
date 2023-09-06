package tokenInterfaces

import "github.com/golang-jwt/jwt/v5"

type ITokenService interface{
	GenerateSecretKey()
	GetTokens()(string, string)
	GenerateTokens(userId string) (string, string, error)
	Valid(tokenString string) bool
	GetClaims(tokenString string) jwt.MapClaims
}