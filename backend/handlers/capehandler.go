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

func GetOldCapes(w http.ResponseWriter, r *http.Request) {

	response, err := controllers.GetOldCapes()

	if err != nil {
		log.Println(err)
		util.ErrorResponse(w, r, "Unknown error")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func GetCapesFull(w http.ResponseWriter, r *http.Request) {
	if !util.IsValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	response, err := json.Marshal(controllers.GetAllCapesFull())

	if err != nil {
		log.Println(err)
		util.ErrorResponse(w, r, "Unknown error")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func GetGapes(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(controllers.GetCapes())

	if err != nil {
		log.Println(err)
		util.ErrorResponse(w, r, "Unknown error")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func DeleteCape(w http.ResponseWriter, r *http.Request) {

	check, perms := util.FullCheck(w, r)

	if !check {
		return
	}

	uuidD, _ := uuid.Parse(mux.Vars(r)["uuid"])

	response, err := controllers.DeleteCape(uuidD)

	if err != nil {
		util.ErrorResponse(w, r, "Bad UUID")
		return
	}
	log.Println(perms.Username + "/" + perms.ID.String() + " has deleted cape " + uuidD.String())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func GetSingleCape(w http.ResponseWriter, r *http.Request) {

	uuid, _ := uuid.Parse(mux.Vars(r)["uuid"])

	response, err := controllers.GetSingleCape(uuid)

	if err != nil {
		util.ErrorResponse(w, r, "No cape with uuid")
		return
	}
	json, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)

}

func SetType(w http.ResponseWriter, r *http.Request) {
	check, _ := util.FullCheck(w, r)

	if !check {
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var cape *models.Cape
	err = json.Unmarshal(b, &cape)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = controllers.SetType(cape)

	if err != nil {
		util.ErrorResponse(w, r, "No cape")
		return
	}

	response, err := json.Marshal(cape)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func SetEnabled(w http.ResponseWriter, r *http.Request) {
	perms := util.AccountCheck(w, r)
	if perms == nil {
		return
	}

	uuid, _ := uuid.Parse(mux.Vars(r)["uuid"])

	err := controllers.SetCapeEnabled(perms, uuid)

	if err != nil {
		util.ErrorResponse(w, r, util.ErrAccess.Error())
		return
	}

	response, _ := json.Marshal(models.Cape{Uuid: uuid})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func GetUsersCapes(w http.ResponseWriter, r *http.Request) {

	uuid, _ := uuid.Parse(mux.Vars(r)["uuid"])

	capes := controllers.GetUsersCapes(uuid)

	if len(capes) == 0 {
		util.ErrorResponse(w, r, "User has no capes")
		return
	}
	response, _ := json.Marshal(capes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}
