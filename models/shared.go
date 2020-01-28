package models

import "time"

type Shared struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"-"`
}
