package startup

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// test if user is authenticated / has running session
		fmt.Println("auth stuff")

		f(w, r)
	}
}
