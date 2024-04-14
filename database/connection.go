package database

import (
	"fmt"
	"os"
	"time"
	"urlshortener/types" 
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
)

// manager having all the variable for database connection
type manager struct{
	connection 	*mongo.Client
	ctx 		context.Context
	cancel 		context.CancelFunc
}

//storing manager in a var
var Mgr Manager 

// declaring Manager
type Manager interface {
	//insert func for inserting data operation
	Insert(interface{}, string) (interface{}, error)
	// after helper function to check for same url existence
	GetUrlFromCode(string, string) (types.UrlDB, error)
	// move to dbcalls
}

// function to connect to database
func ConnectDb(){
	//db host will come from env variable
	uri := os.Getenv("DB_HOST")
	// mongodb client
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri)))
	if err != nil{
		// never use panic in production server as it will break the code and will have to restart
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil{
		panic(err)
	}

	// if everything works ping the database to see if working
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	Mgr = &manager{connection: client, ctx: ctx, cancel: cancel}
}