package services

import (
	"github.com/GDG-KHU-Side/backend-side-project/models"
)

type ReservationService struct {
	db Database
}

func NewReservationService(db Database) *ReservationService {
	return &ReservationService{
		db: db,
	}
}

func (s *ReservationService) GetAllRestaurants(restaurantID int64) ([]models.Reservation, error) {
	return s.db.GetReservations(restaurantID)
}
