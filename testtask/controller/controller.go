package controller

import (
	"testtask2/controller/userreqs"
	"testtask2/controller/userreqs/interfaces"

	"github.com/labstack/echo/v4"
	// "testtask2/model"
)

func Controller() {
	e := echo.New()
	
	

	var userReqs interfaces.IUserReqs = userreqs.UserReqs{Ee: e}

	userReqs.GetRefreshToken() // TODO исправить подключение/отключение БД на каждое обращение по маршруту GET /getToken
	userReqs.PostGetToken()
	
	e.Start(":8080")
}