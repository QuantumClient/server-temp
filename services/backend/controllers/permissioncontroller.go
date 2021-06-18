package controllers

import (
	"backend/db"
	"backend/models"
	"github.com/google/uuid"
	"log"
)

func GetAllAccounts() []models.Permission {
	var (
		perm  models.Permission
		perms []models.Permission
	)

	rows, err := db.Db.Query("SELECT 'uuid', 'admin', 'access', 'createdAt' FROM permissions")

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		rows.Scan(&perm.ID, &perm.Admin, &perm.Access, &perm.CreatedAt)
		res, err := db.Db.Query("SELECT username FROM users WHERE uuid=?", perm.ID)
		if err != nil {
			log.Println(err)
		}
		if res.Next() {
			err := res.Scan(&perm.Username)
			if err != nil {
				log.Fatal(err)
			}
		}
		perms = append(perms, perm)
	}

	defer rows.Close()
	return perms
}

func SetHwid(user *models.ReUser) error {

	_, err := db.Db.Exec("UPDATE permissions SET hwid=? WHERE uuid=?", user.Hwid, user.Uuid)
	if err != nil {
		log.Println(err)
	}

	return err
}

func ResetHwid(uuid uuid.UUID) error {

	_, err := db.Db.Exec("UPDATE permissions SET hwid=NULL WHERE uuid=?", uuid)
	if err != nil {
		log.Println(err)
	}

	return err
}

func SetAdmin(perms *models.Permission, admin bool) error {

	_, err := db.Db.Exec("UPDATE permissions SET admin=? WHERE uuid=?", admin, perms.ID)
	if err != nil {
		log.Println(err)
	}

	return err
}

func SetAccess(perms *models.Permission, access bool) error {

	_, err := db.Db.Exec("UPDATE permissions SET access=? WHERE uuid=?", access, perms.ID)
	if err != nil {
		log.Println(err)
	}

	return err
}
