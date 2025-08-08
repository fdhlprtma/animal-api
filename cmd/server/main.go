package main

import (
	"log"
	"net/http"

	"animal-api/internal/config"
	"animal-api/internal/handler"
	"animal-api/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDB()
	defer db.Close()

	r := mux.NewRouter()

	// Middleware CORS global
	r.Use(middleware.CORSMiddleware)

	// Route login (tanpa JWTAuth)
	r.HandleFunc("/api/login", handler.LoginHandler(db)).Methods("POST")

	// Subrouter yang butuh JWT
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuth)

	// Handler OPTIONS (preflight)
	api.HandleFunc("/animals", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")

	// Rute API yang butuh token
	api.HandleFunc("/animals", handler.GetAnimals(db)).Methods("GET")
	api.HandleFunc("/animals/{id}", handler.GetAnimalByID(db)).Methods("GET")
	api.HandleFunc("/search", handler.SearchAnimals(db)).Methods("GET")
	api.HandleFunc("/animals", handler.CreateAnimal(db)).Methods("POST")

	// Static file handler
	r.PathPrefix("/animal-api/uploads/").Handler(
		http.StripPrefix("/animal-api/uploads/", http.FileServer(http.Dir("./uploads"))),
	)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
