package config

import (
	"context"
//	"os"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


//context
var ctx = func() context.Context {
	return context.Background()
//	return context.WithTimeout(context.Background(), 10*time.Second)
}()


//var dbUserName = os.Getenv("DB_USERNAME")
//var dbPassword = os.Getenv("DB_PASSWORD")

//connect database

func Connect() (*mongo.Database, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://:@cluster0.4xaod.mongodb.net/gomongo"))
	if err != nil {
		return nil, err
	}
	
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("connect database successfully")
	
	return client.Database("gomongo"), nil
}