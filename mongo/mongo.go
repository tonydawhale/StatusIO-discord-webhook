package mongo

import (
	"StatusIO-discord-webhook/types"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func Init() {
	c, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(os.Getenv("MONGO_URI")),
	)

	if err != nil {
		log.Fatal(err)
	}

	MongoClient = c

	if err := MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("Successfully connected to MongoDB")
}

func GetMongoClient() *mongo.Client {
	return MongoClient
}

func GetIncidentFromMongo(id string) (types.MongoIncident, error) {
	database := MongoClient.Database(os.Getenv("MONGO_DATABASE"))
	collection := database.Collection(os.Getenv("MONGO_COLLECTION"))

	var incident types.MongoIncident
	err := collection.FindOne(
		context.TODO(),
		bson.M{"incident_id": id},
	).Decode(&incident)

	return incident, err
}

func SetIncidentData(id string, data types.MongoIncident) (error) {
	database := MongoClient.Database(os.Getenv("MONGO_DATABASE"))
	collection := database.Collection(os.Getenv("MONGO_COLLECTION"))

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{
			"incident_id": id,
		},
		bson.M{
			"$setOnInsert": bson.M{
				"_id": primitive.NewObjectID(),
			},
			"$set": data,
		},
		opts,
	)

	return err
}