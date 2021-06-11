package controllers

import (
	"backend/db"
	"backend/models"
	"backend/util"
	"github.com/google/uuid"
	"log"
)

func Login(user *models.User) (string, error, uuid.UUID) {

	result := GetUserfromName(user.Username)
	if result == nil {
		return "", util.ErrNoAccount, uuid.Nil
	}

	if (result.VerifyPassword(user.Password)) {
		return "", util.ErrBadPassword, uuid.Nil
	}

	token, err := result.GenerateJWT()

	return token, err, result.Uuid


}

func GetUserfromName(name string) *models.User {
	rows, err := db.Db.Query("SELECT * FROM users WHERE username = ?", name)
	if err != nil {
		log.Println(err)
	}
	user := &models.User{}
	for rows.Next() {
		rows.Scan(&user.Uuid,&user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	}
	return user
}

func GetUserfromUUID(uuid uuid.UUID) *models.User {
	rows, err := db.Db.Query("SELECT uuid, username FROM users WHERE uuid = ?", uuid)
	if err != nil {
		log.Println(err)
	}
	user := &models.User{}
	if rows.Next() {
		rows.Scan(&user.Uuid,&user.Username)
	}
	return user
}

func Signup(user *models.User) (string, error) {
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
	_, err = st.Exec(user.Uuid, user.Username)

	if err != nil {
		log.Println(err)
	}


	return user.GenerateJWT()

}
