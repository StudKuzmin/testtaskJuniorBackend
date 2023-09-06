package userreqs

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"testtask2/model"

)

type UserResponse struct {
	Login string `json:"login"`
    Password string `json:"password"`
}

func GetRefreshToken(c echo.Context) error {
    inputGuid := c.QueryParam("guid")

    md := model.ModelService{}
	md.ConnectToDB() // TODO исправить подключение/отключение БД на каждое обращение по маршруту GET /getToken
	defer md.DisconnectFromDB() // TODO исправить подключение/отключение БД на каждое обращение по маршруту GET /getToken

    errorJsonString, err := md.CheckGuid(inputGuid) // Смотрим, есть ли в БД документ с таким guid

    if err==nil{
        tokenJsonString := md.GenerateTokens(inputGuid)
        return c.String(http.StatusOK, tokenJsonString) 
    } else {
        return c.String(http.StatusOK, errorJsonString)
    }
}

func PostGetToken(c echo.Context) error {
    ures := new(UserResponse)
    if err := c.Bind(ures); err != nil {
        return err
    }

    return c.String(http.StatusOK, "POST OK: " + ures.Login + " " + ures.Password)
}