package database

import (
	"context"
	"urlshortener/constant"
	"urlshortener/types"
	"gopkg.in/mgo.v2/bson"
)

// calling everything from manager
func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {
	// connectivity to database and insert function
	// create database connection
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	result, err := inst.InsertOne(context.TODO(), data)
	return result.InsertedID, err
}

// after connection.go helper
func (mgr *manager) GetUrlFromCode(code string, collectionName string) (resp types.UrlDB,err error){
	// making isntance for db and passing query in it
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	err = inst.FindOne(context.TODO(), bson.M{"url_code": code}).Decode(&resp)
	return resp, err
}