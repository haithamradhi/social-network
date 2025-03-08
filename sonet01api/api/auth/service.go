package auth

import (
	"errors"
	"net/http"
	types "sonet01api/api/auth/types"
	"sonet01api/api/models"
	util "sonet01api/util"
	"time"
)

type AuthService struct {
	authRepo AuthRepo
}

type AuthServiceInterface interface {
	CreateUser(req *http.Request) (types.UserResponse, error)
}

// Auth is a simple health check handler
func (h AuthService) CreateUser(req map[string]interface{}) (types.UserResponse, error) {
	// get repo instance & call database ping method
	repo := h.authRepo

	// Validate form fields
	email := req["email"]
	passwordHash := req["password"]
	firstName := req["first_name"]
	lastName := req["last_name"]
	dateOfBirth := req["date_of_birth"]
	avatarPath := req["avatar_path"]
	nickname := req["nickname"]
	aboutMe := req["about_me"]
	isPublic := req["is_public"]

	if email == "" {
		return types.UserResponse{}, errors.New("missing required form field: email")
	}
	if passwordHash == "" {
		return types.UserResponse{}, errors.New("missing required form field: password")
	}
	if firstName == "" {
		return types.UserResponse{}, errors.New("missing required form field: first name")
	}
	if lastName == "" {
		return types.UserResponse{}, errors.New("missing required form field: last name")
	}
	if dateOfBirth == "" {
		return types.UserResponse{}, errors.New("missing required form field: date of birth")
	}
	if avatarPath == "" {
		return types.UserResponse{}, errors.New("missing required form field: avatar path")
	}
	if nickname == "" {
		return types.UserResponse{}, errors.New("missing required form field: nickname")
	}
	if aboutMe == "" {
		return types.UserResponse{}, errors.New("missing required form field: about me")
	}
	if isPublic == "" {
		return types.UserResponse{}, errors.New("missing required form field: is public")
	}

	user := models.UserEntity{
		Email:        req["email"].(string),
		PasswordHash: req["password"].(string),
		FirstName:    req["first_name"].(string),
		LastName:     req["last_name"].(string),
		DateOfBirth:  util.ParseDate(req["date_of_birth"].(string)),
		AvatarPath:   "null",
		Nickname:     req["nickname"].(string),
		AboutMe:      req["about_me"].(string),
		IsPublic:     req["is_public"].(string) == "true",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	response, err := repo.CreateUser(user)
	if err != nil {
		return types.UserResponse{}, err
	}

	return response, nil
}
