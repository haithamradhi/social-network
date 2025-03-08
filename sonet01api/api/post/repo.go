package post

import (
	"sonet01api/api/models"
	db "sonet01api/server/database"
	"strconv"
)

// create repo interface and struct
type PostRepoInterface interface {
	CreatePost(models.UserEntity) (models.PostEntity, error)
}

type PostRepo struct{}

// create function to create a new user
func (h PostRepo) CreatePost(post models.PostEntity) (res models.PostEntity, err error) {

	// create a new user
	// insert into database

	sqlRes, err := db.Exec("INSERT INTO posts (user_id, group_id, content, image_path, privacy_level, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		post.UserId,
		post.GroupId,
		post.Content,
		post.ImagePath,
		post.PrivacyLevel,
		post.CreatedAt,
		post.UpdatedAt,
	)

	if err != nil {
		return models.PostEntity{}, err
	}

	id, err := sqlRes.LastInsertId()
	if err != nil {
		return models.PostEntity{}, err
	}

	res = models.PostEntity{
		Id:           strconv.FormatInt(id, 10),
		UserId:       post.UserId,
		GroupId:      post.GroupId,
		Content:      post.Content,
		ImagePath:    post.ImagePath,
		PrivacyLevel: post.PrivacyLevel,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
	}

	return res, nil

}
