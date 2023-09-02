package interfaces

type IUserReqs interface {
	GetRefreshToken() // GET выполняет Refresh операцию на пару Access, Refresh токенов
	PostGetToken()    //  POST выдает пару Access, Refresh токенов для пользователя сидентификатором (GUID) указанным в параметре запроса
}
