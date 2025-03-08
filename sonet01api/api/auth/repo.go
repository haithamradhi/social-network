package auth

import (
	types "sonet01api/api/auth/types"
	"sonet01api/api/models"
	db "sonet01api/server/database"

	"time"
)

// create repo interface and struct
type AuthRepoInterface interface {
	CreateUser(models.UserEntity) (types.UserResponse, error)
}

type AuthRepo struct{}

// create function to create a new user
func (h AuthRepo) CreateUser(user models.UserEntity) (res types.UserResponse, err error) {

	// create a new user
	// insert into database

	sqlRes, err := db.Exec("INSERT INTO users (email, password_hash, first_name, last_name, date_of_birth, avatar_path, nickname, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NULL, ?, ?, ?)",
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.DateOfBirth,
		user.Nickname,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return types.UserResponse{}, err
	}

	id, err := sqlRes.LastInsertId()
	if err != nil {
		return types.UserResponse{}, err
	}

	_, err = db.Exec("UPDATE users SET avatar_path = ? WHERE id = ?", user.AvatarPath, id)
	if err != nil {
		return types.UserResponse{}, err
	}

	res = types.UserResponse{
		ID:          int(id),
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DateOfBirth: user.DateOfBirth,
		AvatarPath:  user.AvatarPath,
		Nickname:    user.Nickname,
	}

	return res, nil

}
