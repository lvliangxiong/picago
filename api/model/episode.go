package model

import "time"

type Episode struct {
	Id        string    `json:"_id"`
	Title     string    `json:"title"`
	Order     int       `json:"order"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ComicImage struct {
	Id    string `json:"_id"`
	Thumb Thumb  `json:"media"`
}
