package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	. "github.com/tanaka/podlist-backend/config"
	. "github.com/tanaka/podlist-backend/dao"

	"github.com/tanaka/podlist-backend/controllers"
)

var config = Config{}
var dao = DatabaseDAO{}

// Define HTTP request routes
func main() {
	// loads .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get port number
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Print(os.Getenv("PORT"))

	r := mux.NewRouter()

	r.HandleFunc("/", Index).Methods("GET")

	// Posts
	r.HandleFunc("/posts", controllers.AllPostsEndPoint).Methods("GET")
	r.HandleFunc("/post", controllers.CreatePostEndPoint).Methods("POST")
	r.HandleFunc("/post", controllers.UpdatePostEndPoint).Methods("PUT")
	r.HandleFunc("/post", controllers.DeletePostEndPoint).Methods("DELETE")
	r.HandleFunc("/post/{id}", controllers.FindPostEndPoint).Methods("GET")

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Index - to public API
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
