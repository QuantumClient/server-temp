package controllers

import (
	"backend/db"
	"backend/models"
	"backend/util"
	"github.com/google/uuid"
	"log"
)

func Login(user *models.User) (*models.UserResponse, error) {

	result := GetUserfromName(user.Username)
	if result == nil {
		return nil, util.ErrNoAccount
	}

	if result.VerifyPassword(user.Password) {
		return nil, util.ErrBadPassword
	}

	response := &models.UserResponse{
		Uuid:     result.Uuid,
		Username: result.Username,
	}

	acc := models.PermsfromUser(result)
	response.AccessToken = GetJWT(acc)
	response.RefreshToken = GetRefreshToken(acc.ID, result.Password)

	return response, nil

}

func GetUserfromName(name string) *models.User {
	rows, err := db.Db.Query("SELECT * FROM users WHERE username = ?", name)
	if err != nil {
		log.Println(err)
	}
	user := &models.User{}
	for rows.Next() {
		rows.Scan(&user.Uuid, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	}
	defer rows.Close()
	return user
}

func GetUserfromUUID(uuid uuid.UUID) *models.User {
	rows, err := db.Db.Query("SELECT uuid, username FROM users WHERE uuid = ?", uuid)
	if err != nil {
		log.Println(err)
	}
	user := &models.User{}
	if rows.Next() {
		rows.Scan(&user.Uuid, &user.Username)
	}
	defer rows.Close()
	return user
}

func Signup(user *models.User) (string, error) {

	if !util.Alphanumeric3p(user.Username) {
		return "", util.ErrIUsername
	}

	nameCheck := GetUserfromName(user.Username)

	if nameCheck.Uuid != uuid.Nil {
		return "", util.ErrUsername
	}

	stmt, err := db.Db.Prepare("INSERT into users SET uuid=?, username=?, password=?, created_at=?, updated_at=?")
	if err != nil {
		log.Println(err)
	}
	_, queryError := stmt.Exec(user.Uuid, user.Username, user.Password, user.UpdatedAt, user.CreatedAt)

	if queryError != nil {
		log.Println(err)
	}
	st, err := db.Db.Prepare("INSERT into permissions SET uuid=?")
	if err != nil {
		log.Println(err)
	}
	_, err = st.Exec(user.Uuid)

	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	defer st.Close()
	return GetJWT(models.PermsfromUser(user)), nil

}

func getHashedPassword(uuid uuid.UUID) string {
	rows, err := db.Db.Query("SELECT password FROM users WHERE uuid = ?", uuid)
	if err != nil {
		log.Println(err)
	}
	var password string
	if rows.Next() {
		rows.Scan(&password)
	}
	defer rows.Close()
	return password
}
