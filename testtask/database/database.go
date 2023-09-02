package database

import (
	// "context"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"

	"testtask2/database/dbservice"
	"testtask2/database/dbservice/interfaces"
)

func Database(inputGuid string) bool{
	var dbsrvc interfaces.IDBservice
	dbsrvc = &dbservice.DBservice{}
	
	dbsrvc.Connect()
	euser := dbsrvc.FindOne(inputGuid)
	if !euser.Id.IsZero() {
		return true
	}

	dbsrvc.Disconnect()

	return false
	// const uri = "mongodb://localhost:27017"

	// ctx := context.TODO() 
	// opts := options.Client().ApplyURI(uri) 
	// client, err := mongo.Connect(ctx, opts) 
	// if err != nil { 
	// 	panic(err) 
	// } 
	// defer client.Disconnect(ctx)
	// if err := client.Ping(ctx, readpref.Primary()); err != nil { 
	// 	panic(err) 
	// } 

	// //////////////////////
	// coll := client.Database("newdb").Collection("users")
	// filter := bson.D{{"guid", inputGuid}}

	// var euser entities.EUser
	// err = coll.FindOne(context.TODO(), filter).Decode(&euser)

	// if err != nil && err == mongo.ErrNoDocuments {
	// 	println("DOCUMENT NOT FOUND")
	// 	// This error means your query did not match any documents.
	// 	return false
	// }

	// print("OUTPUT: "); println(euser.Guid + " " + euser.Id.String() + " " + euser.Name)

	// return true
}