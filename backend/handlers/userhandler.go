package handlers

import (
	"backend/controllers"
	"backend/models"
	"backend/util"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
		if err == util.ErrNoAccount || err == util.ErrBadPassword {
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
		if err == util.ErrUsername {
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
		log.Println(err)
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

func LinkMCAccount(w http.ResponseWriter, r *http.Request) {
	perms := util.AccountCheck(w, r)
	if perms == nil {
		return
	}
	uuid, _ := uuid.Parse(mux.Vars(r)["uuid"])

	if uuid != perms.ID {
		util.ErrorResponse(w, r, util.ErrAccess.Error())
		return
	}
	key := r.URL.Query().Get("key")

	if len(key) < 8 {
		util.ErrorResponse(w, r, "Invalid key")

		return
	}

	capes, err := controllers.LinkMCAccount(perms, key)
	if err != nil {
		util.ErrorResponse(w, r, "Server side error")

		return
	}

	response, _ := json.Marshal(capes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}
