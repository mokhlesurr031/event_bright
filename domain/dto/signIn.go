package dto

import "time"

type JWTToken struct {
	ExpiredIn time.Duration
	MaxAge    int
	Secret    string
	Message   string
	User      *LoggerInUserData
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoggerInUserData struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
