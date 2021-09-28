package handlers

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetOnline(w http.ResponseWriter, r *http.Request) {

	response, err := json.Marshal(controllers.GetAllOnline())

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func AddToOnline(w http.ResponseWriter, r *http.Request) {
	token, err := controllers.GetToken(r)
	if err != nil || !token.Valid || !token.Claims.(*controllers.JwtCustomClaims).Access {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var online *models.Online
	err = json.Unmarshal(b, &online)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	array := controllers.AddToOnline(online)

	response, err := json.Marshal(array)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func RemoveFromOnline(w http.ResponseWriter, r *http.Request) {

	token, err := controllers.GetToken(r)
	if err != nil || !token.Valid || !token.Claims.(*controllers.JwtCustomClaims).Access {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var online *models.Online
	err = json.Unmarshal(b, &online)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	array := controllers.RemoveFromOnline(online)

	response, err := json.Marshal(array)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}
