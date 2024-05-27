package model

import "time"

type Session struct {
	SessId    string    `json:"sessId"`
	UserId    int       `json:"userId"`
	IsAdmin   bool      `json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt"`
}
