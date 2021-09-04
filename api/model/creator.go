package model

type Creator struct {
	Id         string   `json:"_id"`
	Gender     string   `json:"gender"`
	Name       string   `json:"name"`
	Verified   bool     `json:"verified"`
	Exp        int      `json:"exp"`
	Level      int      `json:"level"`
	Characters []string `json:"characters"`
	Role       string   `json:"role"`
	Title      string   `json:"title"`
	Slogan     string   `json:"slogan"`
	Character  string   `json:"character"`
}
