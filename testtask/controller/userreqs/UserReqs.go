package userreqs

import (
	"github.com/labstack/echo/v4"
)

type UserReqs struct {
	Ee *echo.Echo
}

func (userReqs UserReqs) GetRefreshToken() {
	userReqs.Ee.GET("/get", GetRefreshToken)
}
func (userReqs UserReqs) PostGetToken() {
	userReqs.Ee.POST("/post", PostGetToken)
}
