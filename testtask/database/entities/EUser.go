package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type EUser struct{
	Id primitive.ObjectID `bson:"_id"`
	Name string `bson:"name"`
	Guid string `bson:"guid"`
}