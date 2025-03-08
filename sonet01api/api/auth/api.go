package auth

import (
	"fmt"
	"net/http"
	util "sonet01api/util"
)

type AuthAPI struct {
	authService AuthService
}

type AuthAPIInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

// Create user
func (a AuthAPI) CreateUser(w http.ResponseWriter, r *http.Request) {

	authService := a.authService
	body, err := util.ParseBodyToJSON(r.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	response, err := authService.CreateUser(body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	util.WriteJSON(w, response)
}
