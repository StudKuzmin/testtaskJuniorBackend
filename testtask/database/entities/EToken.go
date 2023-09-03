package entities

type EToken struct{
	accessToken string
	refreshToken string
}
func(et *EToken) GetTokens() (string, string) {
	return et.accessToken, et.refreshToken
}