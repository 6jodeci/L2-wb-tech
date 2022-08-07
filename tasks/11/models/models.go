package models

import "time"

type Event struct {
	ID          string    `json:"id"`
	Header      string    `json:"header"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	CreatorID   string    `json:"creator_id"`
}
