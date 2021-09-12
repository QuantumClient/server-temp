package handlers

import (
	"backend/controllers"
	"backend/models"
	"backend/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {

	response, err := json.Marshal(controllers.GetProjects())

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func UpdateProjectVersion(w http.ResponseWriter, r *http.Request) {
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
	var project *models.Project
	err = json.Unmarshal(b, &project)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = controllers.UpdateProjectVersion(project)

	if err != nil {
		util.ErrorResponse(w, r, "No such project")
		return
	}

	response, err := json.Marshal(project)

	if err != nil {
		log.Println(err)
	}
	log.Println(perms.Username + "/" + perms.ID.String() + " has updated the version for project " + project.Name + " to " + project.Verison)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func GetProject(w http.ResponseWriter, r *http.Request) {
	project := &models.Project{}

	project.Name, _ = mux.Vars(r)["project"]

	response, err := controllers.GetProject(project)

	if err != nil {
		log.Println(err)
	}

	if response == nil {
		util.ErrorResponse(w, r, "No Project with name")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func UpdateProjectLink(w http.ResponseWriter, r *http.Request) {
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
	var project *models.Project
	err = json.Unmarshal(b, &project)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = controllers.UpdateProjectLink(project)

	if err != nil {
		util.ErrorResponse(w, r, "No such project")
		return
	}

	response, err := json.Marshal(project)

	if err != nil {
		log.Println(err)
	}

	log.Println(perms.Username + "/" + perms.ID.String() + " has updated the link for project " + project.Name + " to " + project.Link.String)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}
