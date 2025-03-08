package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBodyToJSON(body io.ReadCloser) (map[string]interface{}, error) {
	var data map[string]interface{}

	var _body []byte
	_body, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(_body, &data)
	return data, nil
}

// func that takes response writer and writes the passed struct as json
func WriteJSON(w http.ResponseWriter, data interface{}) error {
	enc := json.NewEncoder(w)
	w.WriteHeader(http.StatusCreated)
	return enc.Encode(data)
}
