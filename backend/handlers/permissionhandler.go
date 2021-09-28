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
	if !util.IsValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	uuidD := mux.Vars(r)["uuid"]

	if !util.IsUUID(uuidD) {
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
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		util.ErrorResponse(w, r, err.Error())
		return
	}
	var bodyUser *models.ReUser
	err = json.Unmarshal(b, &bodyUser)
	if err != nil {
		util.ErrorResponse(w, r, err.Error())
		return
	}
	uuidD := mux.Vars(r)["uuid"]

	if !util.IsUUID(uuidD) {
		util.ErrorResponse(w, r, "Bad UUID")
		return
	}

	uuidB, _ := uuid.Parse(uuidD)
	if uuidB != bodyUser.Uuid {
		util.ErrorResponse(w, r, util.ErrNoAccount.Error())
		return
	}

	preresponse, err := controllers.CanRun(bodyUser)
	if err != nil {
		util.ErrorResponse(w, r, err.Error())
		return
	}

	response, err := json.Marshal(preresponse)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	token := util.GetJWT(r)
	claims, err := util.ValidateJWT(token)

	if err != nil {
		log.Println(err)

		util.ErrorResponse(w, r, "Invaild AccessToken")
		return
	}

	if claims.Valid == nil {
		log.Println(err)
		util.ErrorResponse(w, r, "Invaild AccessToken")
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
	log.Println(perms.Username + "/" + perms.ID.String() + " has set admin for user " + user.Username + "/" + account.ID.String() + " to " + strconv.FormatBool(account.Admin))

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

	log.Println(perms.Username + "/" + perms.ID.String() + " has reset hwid for " + user.Username + "/" + user.Uuid.String())

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
	log.Println(perms.Username + "/" + perms.ID.String() + " has set access for user " + user.Username + "/" + account.ID.String() + " to " + strconv.FormatBool(account.Admin))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func GetUserKey(w http.ResponseWriter, r *http.Request) {
	token, err := controllers.GetToken(r)

	if err != nil {
		log.Println(err)
		util.ErrorResponse(w, r, util.ErrToken.Error())
		return
	}

	if !token.Claims.(*controllers.JwtCustomClaims).Access {
		util.ErrorResponse(w, r, util.ErrAccess.Error())
		return
	}
	userUuid, _ := uuid.Parse(mux.Vars(r)["uuid"])
	refreshToken := controllers.RefreshFromUUID(userUuid)

	type rB struct {
		Uuid  uuid.UUID `json:"uuid"`
		Token string    `json:"token"`
	}

	response, _ := json.Marshal(rB{
		Uuid:  userUuid,
		Token: refreshToken,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment; filename=key.qt")

	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}
