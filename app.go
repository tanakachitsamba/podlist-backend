package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	. "github.com/tanaka/uphoria/config"
	. "github.com/tanaka/uphoria/dao"

	"github.com/tanaka/uphoria/controllers"
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

	r := mux.NewRouter()

	r.HandleFunc("/", Index).Methods("GET")

	// songs
	r.HandleFunc("/songs", controllers.AllSongsEndPoint).Methods("GET")
	r.HandleFunc("/song", controllers.CreateSongEndPoint).Methods("POST")
	r.HandleFunc("/song", controllers.UpdateSongEndPoint).Methods("PUT")
	r.HandleFunc("/song", controllers.DeleteSongEndPoint).Methods("DELETE")
	r.HandleFunc("/song/{id}", controllers.FindSongEndPoint).Methods("GET")


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
