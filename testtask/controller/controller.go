package controller

import (
	"testtask2/controller/userreqs"
	"testtask2/controller/userreqs/interfaces"

	"github.com/labstack/echo/v4"
)

func Controller() {
	e := echo.New()
	
	var userReqs interfaces.IUserReqs = userreqs.UserReqs{Ee: e}

	userReqs.GetRefreshToken()
	userReqs.PostGetToken()
	
	e.Start(":8080")
}