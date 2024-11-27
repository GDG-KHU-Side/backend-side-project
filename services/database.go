package services

import "github.com/GDG-KHU-Side/backend-side-project/models"

type Database interface {
	GetReservations(restaurantID int64) ([]models.Reservation, error)
	CreateReservation(reservation *models.Reservation) error
	UpdateReservation(reservation *models.Reservation) error
	GetRestaurants() ([]models.Restaurant, error)
	GetRestaurant(id int64) (*models.Restaurant, error)
	GetUser(id int64) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id int64) error
}
