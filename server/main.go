package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
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

var userCache = make(map[int]User)
var cacheMutex sync.RWMutex

// Creating my Disc structure
type Disc struct {
	ID    int    `json:"_id"`
	Name  string `json:"name"`
	Img   string `json:"img"`
	Price string `json:"price"`
}

// Create a slice of Disc structs
 var discCache = []Disc{
 	{ID: 1, Name: "Crave", Img: "https://images-na.ssl-images-amazon.com/images/I/612lLSPhT1L.jpg", Price: "$20"},
 	{ID: 2, Name: "Destroyer", Img: "https://m.media-amazon.com/images/I/813mU5I8bFL.jpg", Price: "$20"},
 	{ID: 3, Name: "Virus", Img: "https://images-na.ssl-images-amazon.com/images/I/51vEEQrB6yL.jpg", Price: "$20"},
 	{ID: 4, Name: "Leopard3", Img: "https://m.media-amazon.com/images/I/81VBeQuv9wL.jpg", Price: "$18"},
 	{ID: 5, Name: "Resistor", Img: "https://m.media-amazon.com/images/I/51ygHUFkaqL.jpg", Price: "$18"},
 	{ID: 6, Name: "Hatchet", Img: "https://us.ftbpic.com/product-amz/westside-discs-origio-burst-hatchet-fairway-disc-golf-driver-great/51eJmz-1MTL._AC_SR480,480_.jpg", Price: "$18"},
 	{ID: 7, Name: "Maestro", Img: "https://www.discgolfmarket.com/cdn/shop/products/DM_ActivePremium_Maestro_1024x1024_bac409dd-62f1-42ce-85f0-56cb75b11cf2_512x512.jpg?v=1597850564", Price: "$18"},
 	{ID: 8, Name: "MX3", Img: "https://www.pbsports.com/cdn/shop/files/MX3-400Glow-1_400x_e4e0cf1d-e4f7-4068-b360-b8056ed75366.jpg?v=1724871751&width=533", Price: "$18"},
 	{ID: 9, Name: "Roc3", Img: "https://discgolffanatic.com/wp-content/uploads/2020/04/Innova-Champion-plastic-roc3-midrange.webp", Price: "$18"},
 	{ID: 10, Name: "P2", Img: "https://i0.wp.com/discgolffanatic.com/wp-content/uploads/2021/12/Discmania-P2-putter.jpg?resize=640,640&ssl=1", Price: "$18"},
 	{ID: 11, Name: "Roach", Img: "https://i0.wp.com/discgolffanatic.com/wp-content/uploads/2021/11/Discraft-Roach-Putter.webp?resize=640,640&ssl=1", Price: "$18"},
 	{ID: 12, Name: "Sensei", Img: "https://m.media-amazon.com/images/I/61B6634eizL.jpg", Price: "$18"},
 }


func main() {
	// Connect to MongoDB
	connectToMongo()
	defer client.Disconnect(context.Background())

	// Create a new Gorilla Mux router
	mux := mux.NewRouter()

	mux.HandleFunc("/", handleRoot).Methods("GET")

	mux.HandleFunc("/users", createUser).Methods("POST")
	mux.HandleFunc("/users/{id}", getUser).Methods("GET")
	mux.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	mux.HandleFunc("/users", getAllUsers).Methods("GET")
	mux.HandleFunc("/discs", getAllDiscs).Methods("GET")
	mux.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	fmt.Println("Server listening to :12000")
	http.ListenAndServe(":12000", mux)
}

// Connect to MongoDB
func connectToMongo() {
	var err error
	// MongoDB URI (change it as per your setup)
	uri := "mongodb://localhost:27017" // or use your MongoDB URI

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

// Get the discs collection from the MongoDB database
func getCollection() *mongo.Collection {
	return client.Database("local").Collection("discs")
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

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := userCache[id]; !ok {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	delete(userCache, id)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	//id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(user)
	fmt.Println(userCache[id])

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id, ok := strconv.Atoi(mux.Vars(r)["id"])
	//id, ok := strconv.Atoi(r.PathValue("id"))
	if ok != nil {
		http.Error(w, ok.Error(), http.StatusBadRequest)
		return
	}
	
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if user.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	user.Id = id
	userCache[id] = user
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	// Get all users from the cache
	cacheMutex.RLock()
	var users []User
	for _, user := range userCache {
		users = append(users, user)
	}
	cacheMutex.RUnlock()

	// Marshal users to JSON
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Get all discs
func getDiscs(w http.ResponseWriter, _ *http.Request) {
	var discs []Disc
	fmt.Println("00")
	collection := getCollection()
	fmt.Println("0")

	// Fetch all documents from the users collection
	cursor, err := collection.Find(context.TODO(), bson.D{})
	fmt.Println("0.5")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// defer cursor.Close(context.Background())

	fmt.Println("1")
	// Decode the results into the users slice
	for cursor.Next(context.Background()) {
		fmt.Println("2")
		var disc Disc
		if err := cursor.Decode(&disc); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		discs = append(discs, disc)
		fmt.Println("3")
	}

	// Respond with the users list
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(discs)
}

 func getAllDiscs(w http.ResponseWriter, r *http.Request) {
 	// Get all users from the cache
 	cacheMutex.RLock()
 	var Discs []Disc
 	// appends all discCache into Discs
 	Discs = append(Discs, discCache...)
	
 	cacheMutex.RUnlock()

 	// Marshal users to JSON
 	w.Header().Set("Content-Type", "application/json")
 	j, err := json.Marshal(Discs)
 	if err != nil {
 		http.Error(w, err.Error(), http.StatusInternalServerError)
 		return
 	}

 	// Send response
 	w.WriteHeader(http.StatusOK)
 	w.Write(j)
 }

func createUser(w http.ResponseWriter, r *http.Request) {
	
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if user.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}



	cacheMutex.Lock()
	user.Id = len(userCache)+1
	userCache[len(userCache)+1] = user
	cacheMutex.Unlock()

	fmt.Println(user)
	fmt.Println(userCache[len(userCache)])

	w.WriteHeader(http.StatusNoContent)
}