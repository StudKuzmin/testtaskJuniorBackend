package model

import (
	"encoding/json"
	"errors"

	"testtask2/database/dbservice"
	"testtask2/database/dbservice/interfaces"
	"testtask2/database/entities"

	"testtask2/model/tokenJwt/tokenInterfaces"
	"testtask2/model/tokenJwt"
)
type ModelService struct {
	dbservice interfaces.IDBservice
	tokenservice tokenInterfaces.ITokenService
}
func (ms *ModelService) ConnectToDB(){
	ms.dbservice = &dbservice.DBservice{}
	ms.dbservice.Connect()
}
func (ms *ModelService) DisconnectFromDB(){
	ms.dbservice.Disconnect()
}
func (ms *ModelService) CheckGuid(inputGuid string) (string, error){
	euser, _ := ms.dbservice.FindOne(inputGuid)

	if !euser.Id.IsZero() {
		eerror := entities.EError{Message: "SUCCESS", Code: 200}
        eerrorJson, _ := json.Marshal(eerror)
        eerrorJsonString := string(eerrorJson)
		return eerrorJsonString, nil
	} else {
        eerror := entities.EError{Message: "DOCUMENT NOT FOUND", Code: 404}
        eerrorJson, _ := json.Marshal(eerror)
        eerrorJsonString := string(eerrorJson)
        return eerrorJsonString, errors.New("DOCUMENT NOT FOUND")
    }
}
func (ms *ModelService) GenerateTokens(guid string) string {
	ms.tokenservice = &tokenJwt.TokenService{}
	ms.tokenservice.GenerateSecretKey()

	accessTokenString, refreshTokenString, _ := ms.tokenservice.GenerateTokens(guid)
	etoken := entities.EToken{AccessToken: accessTokenString, RefreshToken: refreshTokenString}

	etokenJson, _ := json.Marshal(etoken)
	etokenJsonString := string(etokenJson)

	// claims := ms.tokenservice.GetClaims(refreshTokenString)
	// print("CLAIMS: "); println(claims["exp"].(float64)); println(claims["guid"].(string))
	
	return etokenJsonString
}
