package post

import (
	"fmt"
	"net/http"
	util "sonet01api/util"
)

type PostAPI struct {
	postService PostService
}

type PostAPIInterface interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
}

// Create post
func (p PostAPI) CreatePost(w http.ResponseWriter, r *http.Request) {

	postService := p.postService
	body, err := util.ParseBodyToJSON(r.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	response, err := postService.CreatePost(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	util.WriteJSON(w, response)
}
