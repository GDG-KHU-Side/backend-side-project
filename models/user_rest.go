package models

import (
	"time"
	//"github.com/GDG-KHU-Side/backend-side-project/db"
)

type UserRest struct {
	UserID    int64
	RestID    int64
	CreatedAt time.Time
}
