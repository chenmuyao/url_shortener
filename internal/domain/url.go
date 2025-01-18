package domain

import "time"

type Url struct {
	ID        int64     `json:"id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	Count     int64     `json:"count"`
}
