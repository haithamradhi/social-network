package models

import (
	"time"
)

type PostEntity struct {
	Id           string    `json:"id"`
	UserId       string    `json:"user_id"`
	GroupId      string    `json:"group_id"`
	Content      string    `json:"content"`
	ImagePath    string    `json:"image_path"`
	PrivacyLevel string    `json:"privacy_level"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
