package post

import (
	"errors"
	"net/http"
	"sonet01api/api/models"
	"time"
)

type PostService struct {
	postRepo PostRepo
}

type PostServiceInterface interface {
	CreatePost(req *http.Request) (models.PostEntity, error)
}

// Auth is a simple health check handler
func (p PostService) CreatePost(req map[string]interface{}) (models.PostEntity, error) {
	// get repo instance & call database ping method
	repo := p.postRepo

	// Validate fields
	userID := req["user_id"]
	groupID := req["group_id"]
	content := req["content"]
	imagePath := req["image_path"]
	privacyLevel := req["privacy_level"]

	if userID == "" || userID == nil {
		return models.PostEntity{}, errors.New("missing required form field: user id")
	}
	if groupID == "" || groupID == nil {
		return models.PostEntity{}, errors.New("missing required form field: group id")
	}
	if content == "" || content == nil {
		return models.PostEntity{}, errors.New("missing required form field: content")
	}
	if imagePath == "" || imagePath == nil {
		return models.PostEntity{}, errors.New("missing required form field: image path")
	}
	if privacyLevel == "" || privacyLevel == nil {
		return models.PostEntity{}, errors.New("missing required form field: privacy level")
	}

	post := models.PostEntity{
		UserId:       req["user_id"].(string),
		GroupId:      req["group_id"].(string),
		Content:      req["content"].(string),
		ImagePath:    req["image_path"].(string),
		PrivacyLevel: req["privacy_level"].(string),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	response, err := repo.CreatePost(post)
	if err != nil {
		return models.PostEntity{}, err
	}

	return response, nil
}
