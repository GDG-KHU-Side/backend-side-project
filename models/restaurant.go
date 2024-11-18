package models

import (
	"time"
	//"github.com/GDG-KHU-Side/backend-side-project/db"
)

type Restaurant struct {
	ID          int64
	Name        string
	Address     string
	Description string
	PhoneNum    string
	WaitingTime int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
