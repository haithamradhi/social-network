package auth

import "time"

type UserResponse struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	AvatarPath  string    `json:"avatar_path"`
	Nickname    string    `json:"nickname"`
}
