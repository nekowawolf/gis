package gis

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	MongoString string = os.Getenv("MONGOSTRING")
	database    *mongo.Database
)

func init() {
	db, err := MongoConnect("gis")
	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v\n", err)
		os.Exit(1)
	}
	database = db
}

func MongoConnect(dbname string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		return nil, fmt.Errorf("MongoConnect: %v", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("MongoConnect Ping: %v", err)
	}
	return client.Database(dbname), nil
}

func InsertOneDoc(collection string, doc interface{}) (interface{}, error) {
	collectionRef := database.Collection(collection)
	insertResult, err := collectionRef.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, fmt.Errorf("InsertOneDoc: %v", err)
	}
	return insertResult.InsertedID, nil
}

func InsertGisData(province, district, subDistrict, village string, border [][]float64) (interface{}, error) {
	newData := GisSmt5{
		ID:          primitive.NewObjectID(),
		Province:    province,
		District:    district,
		SubDistrict: subDistrict,
		Village:     village,
		Border:      border,
	}

	insertedID, err := InsertOneDoc("gis_data", newData)
	if err != nil {
		return nil, fmt.Errorf("InsertGisData: %v", err)
	}

	return insertedID, nil
}
