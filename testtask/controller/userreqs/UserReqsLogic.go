package userreqs

import (
	"net/http"
    "encoding/json"

	"github.com/labstack/echo/v4"

    "testtask2/database"
    "testtask2/database/entities"
)

type UserResponse struct {
	Login string `json:"login"`
    Password string `json:"password"`
}

func GetRefreshToken(c echo.Context) error {
    inputGuid := c.QueryParam("guid")

    if database.Database(inputGuid) {
        println("START GENERATE TOKEN")
    } else {
        eerror := entities.EError{Message: "DOCUMENT NOT FOUND", Code: 404}
        eerrorJson, _ := json.Marshal(eerror)
        eerrorJsonString := string(eerrorJson)
        return c.String(http.StatusOK, eerrorJsonString)
    }

    eerror := entities.EError{Message: "Internal Server Error", Code: 500}
    eerrorJson, _ := json.Marshal(eerror)
    eerrorJsonString := string(eerrorJson)
    return c.String(http.StatusOK, eerrorJsonString)
}

func PostGetToken(c echo.Context) error {
    ures := new(UserResponse)
    if err := c.Bind(ures); err != nil {
        return err
    }

    return c.String(http.StatusOK, "POST OK: " + ures.Login + " " + ures.Password)
}