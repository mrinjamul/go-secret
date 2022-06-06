package models

import (
	"time"
)

// Message struct
type Message struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"userid"`
	UserName  string    `json:"username"`
	Hash      string    `json:"hash"`
	Message   string    `json:"message"`
	Deleted   bool      `json:"deleted"`
	DeletedAt time.Time `json:"deleted_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User struct
type User struct {
	Id        int64     `json:"id"`
	UserName  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// URL struct
type URL struct {
	Id   int64  `json:"id"`
	URL  string `json:"url"`
	Hash string `json:"hash"`
}

// Response struct
type Response struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}
