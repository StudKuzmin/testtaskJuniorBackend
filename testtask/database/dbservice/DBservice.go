package dbservice

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"testtask2/database/entities"
)

const uri = "mongodb://localhost:27017"

type DBservice struct {
	ctx context.Context
	client *mongo.Client
	err error
}

func(dbservice *DBservice) Connect() {
	dbservice.ctx = context.TODO() 
	opts := options.Client().ApplyURI(uri) 
	dbservice.client, dbservice.err = mongo.Connect(dbservice.ctx, opts)
	if dbservice.err != nil { 
		panic(dbservice.err) 
	}
	err := dbservice.client.Ping(dbservice.ctx, nil)
	if err == nil{
		println("MongoDB: ping ok..")
	}
}
func(dbservice *DBservice) FindOne(param string) (entities.EUser, error) {
	coll := dbservice.client.Database("newdb").Collection("users")
	filter := bson.D{{"guid", param}}

	var euser entities.EUser
	dbservice.err = coll.FindOne(context.TODO(), filter).Decode(&euser)

	if dbservice.err != nil && dbservice.err == mongo.ErrNoDocuments {
		return *new(entities.EUser), errors.New("DOCUMENT NOT FOUND")
	}

	return euser, nil
}
func(dbservice *DBservice) Disconnect() {
	dbservice.client.Disconnect(dbservice.ctx)
}