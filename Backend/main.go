package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people []Person

func main() {
	client, err := GetMongoClient()

	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		PeopleHandler(w, r, client)
	})
	http.HandleFunc("/heath", healthCheckHandler)
	http.HandleFunc("/javascript-response", JavaScriptResponseHandler)

	defer client.Disconnect(context.Background())
	log.Println("server listening on port 8080")
	err = http.ListenAndServe(":8080", corsHandler(http.DefaultServeMux))

	if err != nil {
		log.Fatal(err)
	}

}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	}
}

func PeopleHandler(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	switch r.Method {
	case http.MethodGet:
		getPeople(w, r, client)
	case http.MethodPost:
		postPerson(w, r, client)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func getPeople(w http.ResponseWriter, _ *http.Request, client *mongo.Client) {
	cursor, err := client.Database("social").Collection("users").Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &people); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func postPerson(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := client.Database("social").Collection("users").InsertOne(context.Background(), person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(result)
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	person.ID = insertedID
	jsonResponse, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)

	err = WriteFile("./Backend/people.json", people)
	if err != nil {
		log.Println("Error writing to file:", err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "http web-server works correct")
}

func JavaScriptResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprint(w, `console.log("Hello from Go server!");`)
}

func GetMongoClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}

	fmt.Println("MongoDB-ga muvaffaqiyatli ulanildi")
	return client, nil
}

func WriteFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
