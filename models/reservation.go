package models

import (
	"time"

	"github.com/GDG-KHU-Side/backend-side-project/db"
)

type Reservation struct {
	ID               int64
	RestID           int64
	Count            int
	CustomerPhoneNum string
	CreatedAt        time.Time
	IsEntry          int8
}

func (r *Reservation) Create() error {
	query := `
        INSERT INTO reservation (rest_id, count, customer_phone_num, is_entry)
        VALUES (?, ?, ?, ?)
    `
	result, err := db.DB.Exec(query, r.RestID, r.Count, r.CustomerPhoneNum, r.IsEntry)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	r.ID = id
	return nil
}
