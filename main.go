package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/GDG-KHU-Side/backend-side-project/config"
	"github.com/GDG-KHU-Side/backend-side-project/db"
	"github.com/GDG-KHU-Side/backend-side-project/handlers"
	"github.com/GDG-KHU-Side/backend-side-project/services"
)

func main() {
	conf := config.GetDBConfig()

	err := db.InitDB(conf)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.DB.Close()

	r := mux.NewRouter()

	restaurantService := &services.RestaurantService{}
	restaurantHandler := handlers.NewRestaurantHandler(restaurantService)

	//page
	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/restaurant-list", handlers.RestaurantListHandler)
	r.HandleFunc("/restaurant/{id:[0-9]+}", handlers.RestaurantDetailHandler)

	//api
	r.HandleFunc("/api/restaurant-list", restaurantHandler.GetRestaurantList).Methods("GET")
	r.HandleFunc("/api/restaurant/{id}", restaurantHandler.GetRestaurant).Methods("GET")
	r.HandleFunc("/api/reservation", restaurantHandler.CreateReservation).Methods("POST")
	r.HandleFunc("/api/reservation/{id}/status", restaurantHandler.UpdateReservationStatus).Methods("PUT")

	log.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
