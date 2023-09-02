package interfaces

import (
	"testtask2/database/entities"
)

type IDBservice interface {
	Connect()
	FindOne(param string) entities.EUser
	Disconnect()
}