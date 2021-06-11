package util

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Error struct {
	Status string `json:"status"`
	Message string `json:"error"`
}

var (
	ErrBadPassword = errors.New("key is invalid")
	ErrNoAccount  = errors.New("key is of invalid type")
	ErrUsername  = errors.New("Username is already used")
)

func ErrorResponse(w http.ResponseWriter, r *http.Request, reson string) {

	response, err := json.Marshal(Error{"failed",reson})

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(response)

}

