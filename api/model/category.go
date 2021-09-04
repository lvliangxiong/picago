package model

type Category struct {
	Title  string `json:"title"`
	Thumb  Thumb  `json:"thumb"`
	IsWeb  bool   `json:"isWeb"`
	Active bool   `json:"active"`
	Link   string `json:"link"`
}
