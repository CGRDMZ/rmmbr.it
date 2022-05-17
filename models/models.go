package models

import "time"

type UrlMap struct {
	Id             uint       `json:"id"`
	ShortUrl       string     `json:"shortUrl"`
	LongUrl        string     `json:"longUrl"`
	VisitedCounter uint64     `json:"visitedCounter"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
}

type User struct {
	Id        uint       `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
