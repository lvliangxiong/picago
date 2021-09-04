package model

import "time"

type Comic struct {
	Id     string `json:"_id"`
	Title  string `json:"title"`
	Author string `json:"author"`

	Thumb Thumb `json:"thumb"`

	Description string `json:"description"`
	ChineseTeam string `json:"chineseTeam"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Finished bool `json:"finished"`

	Categories []string `json:"categories"`
	Tags       []string `json:"tags"`

	LikesCount int `json:"likesCount"`
}

type ComicDetail struct {
	Id     string `json:"_id"`
	Title  string `json:"title"`
	Author string `json:"author"`

	Creator Creator `json:"_creator"`

	Thumb Thumb `json:"thumb"`

	Description string `json:"description"`
	ChineseTeam string `json:"chineseTeam"`

	Categories []string `json:"categories"`
	Tags       []string `json:"tags"`

	PagesCount int `json:"pagesCount"`
	EpsCount   int `json:"epsCount"`

	Finished  bool      `json:"finished"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	AllowDownload bool `json:"allowDownload"`
	AllowComment  bool `json:"allowComment"`

	TotalLikes    int `json:"totalLikes"`
	TotalViews    int `json:"totalViews"`
	ViewsCount    int `json:"viewsCount"`
	LikesCount    int `json:"likesCount"`
	CommentsCount int `json:"commentsCount"`

	IsFavourite bool `json:"isFavourite"`
	IsLiked     bool `json:"isLiked"`
}
