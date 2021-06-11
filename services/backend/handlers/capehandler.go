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
	"strconv"
)

func GetCapes(w http.ResponseWriter, r *http.Request) {

	response, err := controllers.GetCapes()
	if r.URL.Query().Get("form") == "true" {
		capes := controllers.GetCapesForm()
		response, err = json.Marshal(capes)
	}

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


func AddCape(w http.ResponseWriter, r *http.Request)  {

	check, perms := util.FullCheck(w, r)

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
	response, _ := controllers.AddCape(cape)

	log.Println(perms.Username + "/" + perms.ID.String() + " has created cape " + cape.Uuid.String() + " with type " +  strconv.Itoa(cape.CapeType))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func GetSingleCape(w http.ResponseWriter, r *http.Request) {

	cape := &models.Cape{}

	cape.Uuid, _ = uuid.Parse(mux.Vars(r)["uuid"])

	response, err := controllers.GetSingleCape(cape)

	if err != nil {
		log.Println(err)
	}

	if response == nil {
		util.ErrorResponse(w, r, "No cape with uuid")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}


func SetType(w http.ResponseWriter, r *http.Request) {
	check, perms := util.FullCheck(w, r)

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
	log.Println(perms.Username + "/" + perms.ID.String() + " has changed cape type of cape " + cape.Uuid.String() + " to" +  strconv.Itoa(cape.CapeType))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}