package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/Rishavbhusal/mangoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const conectionString = "mongodb+srv://rishavbhusal12:<6PRZIPlahPoaPjnI>@cluster0.hjaeh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const collectionName = "watchlist"

// important
var collection *mongo.Collection

// connect with mongo.db

func init() {
	// client option
	clientOption := options.Client().ApplyURI(conectionString)

	// connect to mangoDb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected Successfully!!!")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("collection reference ready!!")
}

// insert 1 rec
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)
}

// update 1 rec
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified Count:", result.ModifiedCount)
}

// delete 1 rec
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("movie got delete with delete count: ", result.DeletedCount)
}

// delete all rec

func deleteAllMovies() {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The number of movies deleted", deleteResult.DeletedCount)

}

// get all rec

func getAllMovie() {
	// cursor vaneko auta matra data result ma basne jasto hoina collection of data ho jaslai access garnav loop lagaunu parcha
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			f
			log.Fatal(err)
		}
		movies = append(movies, movie)

	}
	defer cursor.Close(context.Background())

}
