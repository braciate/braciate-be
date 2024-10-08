package entity

import "time"

type Assets struct {
	ID           string
	UserID       string
	LkmID        string
	NominationID string
	Url          string
	CreatedAt    time.Time
	updated_at   time.Time
}
