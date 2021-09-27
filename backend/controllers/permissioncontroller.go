package controllers

import (
	"backend/db"
	"backend/models"
	"backend/util"
	"database/sql"
	"github.com/google/uuid"
	"log"
)

func GetAllAccounts() []models.Permission {
	var (
		perm  models.Permission
		perms []models.Permission
	)

	rows, err := db.Db.Query("SELECT users.uuid, username, p.admin, p.access, p.created_at FROM users INNER JOIN permissions p on users.uuid = p.uuid ORDER BY created_at")

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		rows.Scan(&perm.ID, &perm.Username, &perm.Admin, &perm.Access, &perm.CreatedAt)
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

func CanRun(user *models.ReUser) (models.AuthResponse, error) {
	res, err := db.Db.Query("SELECT u.uuid, username, hwid, admin, u.password FROM permissions p JOIN users u ON p.uuid = u.uuid AND p.access = 1 AND u.username = ? AND u.uuid = ?", user.Username, user.Uuid)
	defer res.Close()

	var response models.AuthResponse
	if err != nil {
		log.Println(err)
		return response, util.ErrAccess
	}

	var (
		hwid  sql.NullString
		admin bool
		pass  string
	)

	if res.Next() {
		err = res.Scan(&response.Uuid, &response.Username, &hwid, &admin, &pass)

		if err != nil {
			log.Fatal(err)
		}
	}
	if user.VerifyPassword(pass) {
		return response, util.ErrBadPassword
	}

	if hwid.Valid && hwid.String == user.Hwid {
		response.Status = 1
	} else {
		err = SetHwid(user)
		if err != nil {
			log.Println(err)
		}
	}

	token, err := models.GenerateJWT(user.Uuid.String(), user.Username, admin, true)
	if err != nil {
		log.Println(err)
	}
	response.Token = token

	return response, nil
}

func LinkMCAccount(perms *models.Permission, key string) (*models.Cape, error) {
	i := 2
	if perms.Access {
		i = 1
	}
	_, err := db.Db.Exec("INSERT INTO capes (`key`, account_uuid, type) VALUE (?, ?, ?)", key, perms.ID, i)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	cape := &models.Cape{}
	res, err := db.Db.Query("SELECT vma.player_uuid, vma.player_username, type, url, enabled, account_uuid FROM capes INNER JOIN verified_mc_accounts vma ON capes.`key` = vma.`keys` WHERE `key` = ?", key)
	defer res.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if res.Next() {
		err := res.Scan(&cape.Uuid, &cape.Username, &cape.CapeType, &cape.Url, &cape.Enabled, &cape.Owner_uuid)
		if err != nil {
			log.Println(err)
		}
	}
	return cape, nil

}
