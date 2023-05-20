package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
  "errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Counter struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Value int                `bson:"count"`
}

var count Counter // initialize counter to zero

func connectDB() (*mongo.Client, error) {
	// Connect to MongoDB
  if _, exist := os.LookupEnv("MONGO_HOST"); !exist {
    return nil, errors.New("Please provide MONGO_HOST env variable")
  }
  port := "27017"
  if _, exist := os.LookupEnv("MONGO_PORT"); exist {
    port = os.Getenv("MONGO_PORT")
  }
	clientOptions := options.Client().ApplyURI("mongodb://"+os.Getenv("MONGO_HOST")+":"+port)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check if the connection is successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to mongodb")

	return client, nil
}

func initializeCount() error {
	// Retrieve the count value from the database
	client, err := connectDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("counter-db").Collection("counter")
	filter := bson.M{}
	err = collection.FindOne(context.Background(), filter).Decode(&count)
	if err != nil {
		// If no count value is found, initialize count to zero and insert it into the database
    count = Counter{Value: 0}
		_, err = collection.InsertOne(context.Background(), count)
		if err != nil {
			return err
		}
	}

	return nil
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	// Return the current count value
	json.NewEncoder(w).Encode(count)
}

func increaseHandler(w http.ResponseWriter, r *http.Request) {
	// Increase the count value and return the updated count
	count.Value++
	client, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("counter-db").Collection("counter")
	filter := bson.M{}
	update := bson.M{"$set": bson.M{"count": count.Value}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(count)
}

func decreaseHandler(w http.ResponseWriter, r *http.Request) {
	// Decrease the count value and return the updated count
	count.Value--
	client, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("counter-db").Collection("counter")
	filter := bson.M{}
	update := bson.M{"$set": bson.M{"count": count.Value}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(count)
}

func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		next(w, r)
	}
}

func main() {
	err := initializeCount()
	if err != nil {
		fmt.Println("Failed to initialize count:", err)
		os.Exit(1)
	}

	http.HandleFunc("/count", enableCors(countHandler))
	http.HandleFunc("/increase", enableCors(increaseHandler))
	http.HandleFunc("/decrease", enableCors(decreaseHandler))

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
