package clients

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"home-server/models"
	"log"
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

func WriteSensorData(data *models.SensorDataOut) (result interface{}, err error) {
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
		{"humidity", data.Humidity},
		{"time", primitive.NewDateTimeFromTime(data.Time)}})
	if err != nil {
		return
	}
	result = res.InsertedID
	return
}

func WriteBoilerData(data *models.BoilerDataOut) (result interface{}, err error) {
	client, ctx := getClient()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("test").Collection("boilerData")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{
		{"waterTemperature", data.WaterTemperature},
		{"isRunning", data.IsRunning},
		{"time", primitive.NewDateTimeFromTime(data.Time)}})
	if err != nil {
		return
	}
	result = res.InsertedID
	return
}

func ReadSensorsData(from time.Time, to time.Time) []models.SensorDataOut {
	client, ctx := getClient()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := getCollection(client)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{
		"time": bson.M{"$gte": from, "$lt": to}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var results []models.SensorDataMongo
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for cursor.Next(ctx) {
		var data models.SensorDataMongo
		err := cursor.Decode(&data)
		if err != nil {
			panic(err)
		}
		results = append(results, data)
	}

	var dtoResults []models.SensorDataOut
	for _, data := range results {
		singleRecord := models.SensorDataOut{
			Temperature: data.Temperature,
			Humidity:    data.Humidity,
			Time:        data.Time.Time(),
		}
		dtoResults = append(dtoResults, singleRecord)
	}

	return dtoResults
}

func ReadBoilerData(from time.Time, to time.Time) []models.BoilerDataOut {
	client, ctx := getClient()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("test").Collection("boilerData")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{
		"time": bson.M{"$gte": from, "$lt": to}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var results []models.BoilerDataMongo
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for cursor.Next(ctx) {
		var data models.BoilerDataMongo
		err := cursor.Decode(&data)
		if err != nil {
			panic(err)
		}
		results = append(results, data)
	}

	var dtoResults []models.BoilerDataOut
	for _, data := range results {
		singleRecord := models.BoilerDataOut{
			WaterTemperature: data.WaterTemperature,
			IsRunning:        data.IsRunning,
			Time:             data.Time.Time(),
		}
		dtoResults = append(dtoResults, singleRecord)
	}

	return dtoResults
}

func ReadLatestData() models.AllData {
	client, ctx := getClient()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := getCollection(client)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var lastSensorRecord models.SensorDataMongo
	opts := options.FindOne().SetSort(bson.M{"$natural": -1})
	if err := collection.FindOne(ctx, bson.M{}, opts).Decode(&lastSensorRecord); err != nil {
		log.Fatal(err)
	}

	collection = client.Database("test").Collection("boilerData")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var lastBoilerRecord models.BoilerDataMongo
	opts = options.FindOne().SetSort(bson.M{"$natural": -1})
	if err := collection.FindOne(ctx, bson.M{}, opts).Decode(&lastBoilerRecord); err != nil {
		log.Fatal(err)
	}
	return models.AllData{
		Temperature:           lastSensorRecord.Temperature,
		Humidity:              lastSensorRecord.Humidity,
		BaseLastUpdatedTime:   lastSensorRecord.Time.Time(),
		WaterTemperature:      lastBoilerRecord.WaterTemperature,
		IsRunning:             lastBoilerRecord.IsRunning,
		BoilerLastUpdatedTime: lastBoilerRecord.Time.Time(),
	}
}
