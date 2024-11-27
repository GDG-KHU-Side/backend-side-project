package models

import (
	"time"
)

type Reservation struct {
	ID               int64     `json:"id"`
	RestID           int64     `json:"rest_id"`
	Count            int       `json:"count"`
	CustomerPhoneNum string    `json:"customer_phone_num"`
	CreatedAt        time.Time `json:"created_at"`
	IsEntry          int8      `json:"is_entry"`
}
