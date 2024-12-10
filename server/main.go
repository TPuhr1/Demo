package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "strconv"
	// "sync"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB client
var client *mongo.Client

type User struct {
	Name string `json:"name"`
	Id int `json:"id"`
}

// Creating my Disc structure
type Disc struct {
	ID    int    `json:"_id"`
	Name  string `json:"name"`
	Img   string `json:"img"`
	Price string `json:"price"`
	Type string `json:"type"`
}

func main() {
	// Connect to MongoDB
	connectToMongo()
	defer client.Disconnect(context.Background())

	// Create a new Gorilla Mux router
	mux := mux.NewRouter()

	mux.HandleFunc("/", handleRoot).Methods("GET")

	mux.HandleFunc("/discs", getDiscs).Methods("GET")
	mux.HandleFunc("/discType", getType).Methods("GET")
	fmt.Println("Server listening to :12000")
	http.ListenAndServe(":12000", mux)
}

 // Connect to MongoDB
 func connectToMongo() {
 	var err error
 	// MongoDB URI (change it as per your setup)
 	uri := "mongodb+srv://admin:0qww294e@cluster0.5gxx4.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0" // or use your MongoDB URI

 	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
 	if err != nil {
 		log.Fatal(err)
 	}

 	// Establish connection with a timeout context
 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
 	defer cancel()

 	err = client.Connect(ctx)
 	if err != nil {
 		log.Fatal(err)
 	}

 	fmt.Println("Connected to MongoDB!")
 }

func getCollection() *mongo.Collection {
    // Create a client option and use the new Mongo Connect method
    clientOptions := options.Client().ApplyURI("mongodb+srv://admin:0qww294e@cluster0.5gxx4.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

    // Connect to MongoDB using the new Connect method
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    // Check the connection to MongoDB
    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal("Failed to ping MongoDB:", err)
    }

    // Access and return the collection
    return client.Database("newDB").Collection("discs")
}


func handleCORS(w http.ResponseWriter, r *http.Request) {
	// Allow all origins (*), you can replace this with a specific origin if needed
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// If the method is OPTIONS, return a response with 200 status for preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
}


func handleRoot(w http.ResponseWriter, r *http.Request) {
		// Call the CORS handler to set headers
		handleCORS(w, r)

		// Example response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, this is a CORS-enabled response!")
}

// Get all discs
func getDiscs(w http.ResponseWriter, _ *http.Request) {
	var discs []Disc
	collection := getCollection()

	// Set up a context with a timeout for the query
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	// Fetch all documents from the users collection
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	 defer cursor.Close(ctx)

	// Decode the results into the users slice
	for cursor.Next(ctx) {
		var disc Disc
		if err := cursor.Decode(&disc); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		discs = append(discs, disc)
	}

	// Check if any errors occurred during cursor iteration
    if err := cursor.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return the data as JSON
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(discs); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Get all discs, optionally filtered by query parameters
func getType(w http.ResponseWriter, r *http.Request) {
    var discs []Disc
    collection := getCollection()

    // Extract filter parameters from the query string (e.g., ?genre=Rock)
    Type := r.URL.Query().Get("type")

    // Create a MongoDB filter
    filter := bson.D{}

    // Add filters based on query parameters (only if they are provided)
    if Type != "" {
        filter = append(filter, bson.E{Key: "type", Value: Type})
    }
    

    // Set up a context with a timeout for the query
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Fetch documents from the collection based on the constructed filter
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    // Decode the results into the discs slice
    for cursor.Next(ctx) {
        var disc Disc
        if err := cursor.Decode(&disc); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        discs = append(discs, disc)
    }

    // Handle errors from cursor iteration
    if err := cursor.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return the filtered list of discs as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(discs)
}