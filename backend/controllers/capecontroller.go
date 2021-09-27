package controllers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"github.com/google/uuid"
	"log"
)

func GetOldCapes() ([]byte, error) {

	rows, err := db.Db.Query("SELECT vma.player_uuid, type FROM capes INNER JOIN verified_mc_accounts vma ON capes.`key` = vma.`keys` AND enabled = 1")

	if err != nil {
		log.Println(err)
	}

	capes := make(map[string]int)
	for rows.Next() {
		var uuid string
		var t int

		err = rows.Scan(&uuid, &t)
		if err != nil {
			log.Println(err)
		}

		capes[uuid] = t
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Println(err)
	}

	return json.Marshal(capes)

}

func DeleteCape(uuid uuid.UUID) ([]byte, error) {

	_, err := db.Db.Exec("DELETE capes FROM capes INNER JOIN verified_mc_accounts vma on capes.`key` = vma.`keys` WHERE vma.player_uuid = ?;", uuid)

	if err != nil {
		log.Println(err)
	}

	return json.Marshal(map[string]string{"DELETED": uuid.String()})

}

func GetSingleCape(uuid uuid.UUID) (models.Cape, error) {

	var cape models.Cape
	err := db.Db.QueryRow("SELECT vma.player_uuid, vma.player_username, type, url, enabled, account_uuid FROM capes INNER JOIN verified_mc_accounts vma ON capes.`key` = vma.`keys` WHERE player_uuid = ?", cape.Uuid).Scan(&cape.Uuid, &cape.Username, &cape.CapeType, &cape.Url, &cape.Enabled, &cape.Owner_uuid)
	if err != nil {
		log.Println(err)
	}

	return cape, err
}

func GetUsersCapes(uuid uuid.UUID) []models.Cape {
	var (
		cape  models.Cape
		capes []models.Cape
	)

	rows, err := db.Db.Query("SELECT vma.player_uuid, vma.player_username, type, url, enabled, account_uuid FROM capes INNER JOIN verified_mc_accounts vma ON capes.`key` = vma.`keys` WHERE account_uuid = ?", uuid)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		rows.Scan(&cape.Uuid, &cape.Username, &cape.CapeType, &cape.Url, &cape.Enabled, &cape.Owner_uuid)
		capes = append(capes, cape)
	}

	defer rows.Close()
	return capes
}

func SetType(cape *models.Cape) error {

	_, err := db.Db.Exec("UPDATE capes INNER JOIN verified_mc_accounts vma on capes.`key` = vma.`keys` SET type = ? WHERE vma.player_uuid = ?", cape.CapeType, cape.Uuid)
	if err != nil {
		log.Println(err)
	}

	return err
}

func GetAllCapesFull() []models.Cape {
	var (
		cape  models.Cape
		capes []models.Cape
	)

	rows, err := db.Db.Query("SELECT vma.player_uuid, vma.player_username, type, url, enabled, account_uuid, u.username FROM capes INNER JOIN verified_mc_accounts vma ON capes.`key` = vma.`keys` INNER JOIN users u on account_uuid = u.uuid order by account_uuid")

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		rows.Scan(&cape.Uuid, &cape.Username, &cape.CapeType, &cape.Url, &cape.Enabled, &cape.Owner_uuid, &cape.Owner_username)
		capes = append(capes, cape)
	}

	defer rows.Close()
	return capes
}

func SetCapeEnabled(perm *models.Permission, uuid uuid.UUID) error {
	var err error
	if perm.Admin {
		stmt, _ := db.Db.Prepare("UPDATE capes INNER JOIN verified_mc_accounts vma on capes.`key` = vma.`keys` SET enabled = NOT enabled WHERE vma.player_uuid = ?")
		_, err = stmt.Exec(uuid)
		stmt.Close()
	} else {
		stmt, _ := db.Db.Prepare("UPDATE capes INNER JOIN verified_mc_accounts vma on capes.`key` = vma.`keys` SET enabled = NOT enabled WHERE vma.player_uuid = ? AND account_uuid = ?")
		_, err = stmt.Exec(uuid, perm.ID)
		stmt.Close()
	}

	if err != nil {
		log.Println(err)
	}

	return err
}

func GetCapes() []models.Cape {
	var (
		cape  models.Cape
		capes []models.Cape
	)

	rows, err := db.Db.Query("SELECT vma.player_uuid, vma.player_username, type, url FROM capes INNER JOIN verified_mc_accounts vma ON capes.`key` = vma.`keys` AND enabled = 1")

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		rows.Scan(&cape.Uuid, &cape.Username, &cape.CapeType, &cape.Url)
		capes = append(capes, cape)
	}

	defer rows.Close()
	return capes
}
