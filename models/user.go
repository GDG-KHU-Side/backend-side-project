package models

import (
	"time"
	//"github.com/GDG-KHU-Side/backend-side-project/db"
)

type User struct {
	ID        int64
	Email     string
	Password  string
	PhoneNum  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
