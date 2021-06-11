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

func GetPerms(w http.ResponseWriter, r *http.Request) {

	uuidD := mux.Vars(r)["uuid"]

	if (!util.IsUUID(uuidD)) {
		util.ErrorResponse(w, r, "Bad UUID")
		return
	}

	uuid, err := uuid.Parse(uuidD)
	if err != nil {
		log.Println(err)
	}

	user := controllers.GetUserfromUUID(uuid)

	result := models.PermsfromUser(user)
	if result == nil {
		util.ErrorResponse(w, r, "Unknown Error")
		return
	}

	response, err := json.Marshal(result)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}


func GetAllAccounts(w http.ResponseWriter, r *http.Request) {

	response, err := json.Marshal(controllers.GetAllAccounts())

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func CanRun(w http.ResponseWriter, r *http.Request) {

	//all of this is shit code and need to be redone but i just wanna be done
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var bodyUser *models.ReUser
	err = json.Unmarshal(b, &bodyUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	loginuser := &models.User{
		Uuid:      bodyUser.Uuid,
		Username:  bodyUser.Username,
		Password:  bodyUser.Password,
	}

	_, err, uuidA := controllers.Login(loginuser)

	if err != nil {
		log.Println(err)
		if (err == util.ErrNoAccount || err == util.ErrBadPassword) {
			util.ErrorResponse(w, r, err.Error())
			return
		}
	}

	uuidD := mux.Vars(r)["uuid"]

	if (!util.IsUUID(uuidD)) {
		util.ErrorResponse(w, r, "Bad UUID")
		return
	}

	uuidB, err := uuid.Parse(uuidD)
	if err != nil {
		log.Println(err)
	}
	if uuidB != uuidA {
		util.ErrorResponse(w, r, util.ErrNoAccount.Error())
		return
	}

	userUUID := controllers.GetUserfromUUID(uuidB)

	perms := models.PermsfromUser(userUUID)
	if perms == nil {
		util.ErrorResponse(w, r, "Unknown Error")
		return
	}
	s := 0
	if perms.Access {
		if perms.Hwid.Valid {
			if perms.Hwid.String == bodyUser.Hwid {
				s = 1
			}
		} else {
			controllers.SetHwid(bodyUser)
			s = 1
		}
	}


	preresponse := &authresponse{
		Status:   s,
		Uuid:     uuidA,
		Username: bodyUser.Username,
	}

	response, err := json.Marshal(preresponse)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

type authresponse struct {
	Status int 			`json:"status"`
	Uuid uuid.UUID 		`json:"uuid"`
	Username  string    `json:"username"`
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
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

	response, err := json.Marshal(claims)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func SetAdmin(w http.ResponseWriter, r *http.Request) {
	check, perms := util.FullCheck(w, r)

	if !check {
		return
	}
	uuidD := mux.Vars(r)["uuid"]

	uuid, err := uuid.Parse(uuidD)
	if err != nil {
		log.Println(err)
	}

	user := controllers.GetUserfromUUID(uuid)

	account := models.PermsfromUser(user)

	err = controllers.SetAdmin(account, !account.Admin)

	if err != nil {
		log.Println(err)
	}
	account.Admin = !account.Admin

	response, err := json.Marshal(account)

	if err != nil {
		log.Println(err)
	}
	log.Println(perms.Username + "/" + perms.ID.String() + " has set admin for user " + user.Username + "/" +  account.ID.String() + " to " + strconv.FormatBool(account.Admin))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func ResetHWID(w http.ResponseWriter, r *http.Request) {
	check, perms := util.FullCheck(w, r)

	if !check {
		return
	}
	uuidD := mux.Vars(r)["uuid"]

	uuid, err := uuid.Parse(uuidD)
	if err != nil {
		log.Println(err)
	}

	user := controllers.GetUserfromUUID(uuid)


	err = controllers.ResetHwid(user.Uuid)

	if err != nil {
		log.Println(err)
	}

	response, err := json.Marshal(user)

	if err != nil {
		log.Println(err)
	}

	log.Println(perms.Username + "/" + perms.ID.String() + " has reset hwid for " + user.Username + "/" +  user.Uuid.String())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func SetAccess(w http.ResponseWriter, r *http.Request) {
	check, perms := util.FullCheck(w, r)

	if !check {
		return
	}
	uuidD := mux.Vars(r)["uuid"]

	uuid, err := uuid.Parse(uuidD)
	if err != nil {
		log.Println(err)
	}

	user := controllers.GetUserfromUUID(uuid)

	account := models.PermsfromUser(user)

	err = controllers.SetAccess(account, !account.Access)

	if err != nil {
		log.Println(err)
	}
	account.Access = !account.Access

	response, err := json.Marshal(account)

	if err != nil {
		log.Println(err)
	}
	log.Println(perms.Username + "/" + perms.ID.String() + " has set access for user " + user.Username + "/" +  account.ID.String() + " to " + strconv.FormatBool(account.Admin))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}