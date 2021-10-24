package handlers

import (
	"backend/controllers"
	"backend/models"
	"backend/util"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func GetPerms(w http.ResponseWriter, r *http.Request) {
	token, err := controllers.GetToken(r)
	if err != nil || !token.Valid {
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

func CanRunLeg(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		util.ErrorResponse(w, r, err.Error())
		return
	}
	var bodyUser *models.LegUserCheck
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

	preresponse, err := controllers.CanRunLeg(bodyUser)
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
	token, err := controllers.GetToken(r)
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	response, err := json.Marshal(token.Claims)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func SetAdmin(w http.ResponseWriter, r *http.Request) {
	token, err := controllers.GetToken(r)
	if err != nil || !token.Valid || !token.Claims.(*controllers.JwtCustomClaims).Admin {
		w.WriteHeader(http.StatusUnauthorized)
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
	log.Println(token.Claims.(*controllers.JwtCustomClaims).Username + "/" + token.Claims.(*controllers.JwtCustomClaims).Uuid + " has set admin for user " + user.Username + "/" + account.ID.String() + " to " + strconv.FormatBool(account.Admin))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func ResetHWID(w http.ResponseWriter, r *http.Request) {
	token, err := controllers.GetToken(r)
	if err != nil || !token.Valid || !token.Claims.(*controllers.JwtCustomClaims).Admin {
		w.WriteHeader(http.StatusUnauthorized)
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

	log.Println(token.Claims.(*controllers.JwtCustomClaims).Username + "/" + token.Claims.(*controllers.JwtCustomClaims).Uuid + " has reset hwid for " + user.Username + "/" + user.Uuid.String())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func SetAccess(w http.ResponseWriter, r *http.Request) {
	token, err := controllers.GetToken(r)
	if err != nil || !token.Valid || !token.Claims.(*controllers.JwtCustomClaims).Admin {
		w.WriteHeader(http.StatusUnauthorized)
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
	log.Println(token.Claims.(*controllers.JwtCustomClaims).Username + "/" + token.Claims.(*controllers.JwtCustomClaims).Uuid + " has set access for user " + user.Username + "/" + account.ID.String() + " to " + strconv.FormatBool(account.Admin))

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

	key, err := controllers.GenKey(token)
	if err != nil {
		util.ErrorResponse(w, r, err.Error())
		return
	}
	var response []byte
	if _, ok := r.URL.Query()["json"]; ok {
		response, _ = json.Marshal(key)
		w.Header().Set("Content-Type", "application/json")
	} else {
		response = []byte(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s\n%s\n%s", key.Uuid, key.RefreshToken, key.CheckSum))))
		w.Header().Set("Content-Type", "application/text")

	}

	w.Header().Set("Content-Disposition", "attachment; filename=key.qt")

	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func Verify(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		util.ErrorResponse(w, r, err.Error())
		return
	}
	var userReq *models.AuthUserReq
	err = json.Unmarshal(b, &userReq)
	if err != nil {
		util.ErrorResponse(w, r, err.Error())
		return
	}
	sum := md5.Sum([]byte(userReq.Uuid.String() + userReq.RefreshToken))
	if userReq.CheckSum != hex.EncodeToString(sum[:]) {
		util.ErrorResponse(w, r, util.ErrInvalidSum.Error())
		return
	}
	token, err := controllers.JWTFromRefresh(userReq.RefreshToken)
	if err != nil || !token.Valid {
		util.ErrorResponse(w, r, util.ErrToken.Error())
		return
	}
	claims, ok := token.Claims.(*controllers.JwtRefreshClaims)
	if !ok {
		util.ErrorResponse(w, r, util.ErrToken.Error())
		return
	}
	uuidT, _ := uuid.Parse(claims.Uuid)
	userReq.Uuid = uuidT

	authRe, err := controllers.Verify(userReq)
	if err != nil {
		util.ErrorResponse(w, r, err.Error())
		return
	}

	response, err := json.Marshal(authRe)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}
