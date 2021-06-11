package handlers

import (
	"backend/controllers"
	"backend/models"
	"backend/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var user *models.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	token, err, uuid := controllers.Login(user)

	if err != nil {
		log.Println(err)
		if (err == util.ErrNoAccount || err == util.ErrBadPassword) {
			util.ErrorResponse(w, r, err.Error())
			return
		}
	}

	response := models.UserResponse{
		Uuid:     uuid,
		Username: user.Username,
		Token:    token,
	}

	json, err := json.Marshal(response)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)

}



func Signup(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var bodyUser *models.User
	err = json.Unmarshal(b, &bodyUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user, err := models.NewUser(bodyUser.Username, bodyUser.Password)
	if err != nil {
		log.Println(err)
	}

	token, err := controllers.Signup(user)

	if err != nil {
		log.Println(err)
		if (err == util.ErrUsername) {
			util.ErrorResponse(w, r, err.Error())
			return
		}
	}

	response := models.UserResponse{
		Uuid:     user.Uuid,
		Username: user.Username,
		Token:    token,
	}

	json, err := json.Marshal(response)

	if err != nil {
		log.Println(err)
	}
	log.Println(response.Username + "/" + response.Uuid.String() + " has just made an account")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)

}

func Me(w http.ResponseWriter, r *http.Request) {

	token := util.GetJWT(r)

	claims, err := util.ValidateJWT(token)

	if err != nil {
		log.Println(err )
		util.ErrorResponse(w, r, "Invaild Token")
		return
	}

	if claims.Valid == nil {
		log.Println(err)
		util.ErrorResponse(w, r, "Invaild Token")
		return
	}

	_, perms := util.GetAccountsFromToken(claims)

	response, err := json.Marshal(perms)

	if err != nil {
		log.Println(err)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}