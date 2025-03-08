package startup

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// read body

		// get token from header
		token := r.Header.Get("Authorization")
		fmt.Println(token)

		f(w, r)
	}
}
