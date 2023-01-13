package clients

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"home-server/models"
	"time"
)

func getClient() (*mongo.Client, context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, ctx
	}
	return client, ctx
}

func getCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("test").Collection("sensorData")
	return collection
}

func WriteSensorData(data *models.SensorData) (result interface{}, err error) {
	client, ctx := getClient()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := getCollection(client)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{
		{"temperature", data.Temperature},
		{"humidity", data.Humidity}})
	if err != nil {
		return
	}
	result = res.InsertedID
	return
}

func ReadSensorsData() []models.SensorData {
	client, ctx := getClient()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := getCollection(client)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var results []models.SensorData
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		err := cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
	}
	return results
}
