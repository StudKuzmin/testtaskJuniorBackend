package tokenInterfaces

type ITokenService interface{
	GenerateSecretKey()
	GetTokens()(string, string)
	GenerateToken(userId string) (string, error)
	Valid(tokenString string) bool
}