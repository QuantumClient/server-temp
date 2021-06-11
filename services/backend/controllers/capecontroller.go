package controllers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"github.com/google/uuid"
	"log"
)

func GetCapes() ([]byte, error) {

	rows, err := db.Db.Query("SELECT * FROM capes")

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
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	return json.Marshal(capes)

}


func DeleteCape(uuid uuid.UUID) ([]byte, error) {

	_, err := db.Db.Exec("DELETE FROM capes WHERE uuid = ?", uuid)

	if err != nil {
		log.Println(err)
	}

	return json.Marshal(map[string]string{"DELETED": uuid.String()})


}

func AddCape(cape *models.Cape) ([]byte, error) {


	pre, err := db.Db.Prepare("INSERT INTO capes(uuid, type) VALUES (?, ?)")

	if err != nil {
		log.Println(err)
	}

	_, err = pre.Exec(cape.Uuid, cape.CapeType)

	if err != nil {
		log.Println(err)
	}

	return json.Marshal(map[string]string{"ADDED": cape.Uuid.String(), "TYPE": string(cape.CapeType)})


}

func GetSingleCape(cape *models.Cape) ([]byte, error) {

	res, err := db.Db.Query("SELECT * FROM capes WHERE uuid=?", cape.Uuid)
	defer res.Close()
	if err != nil {
		log.Println(err)
	}

	if res.Next() {
		err := res.Scan(&cape.Uuid, &cape.CapeType)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		return nil, nil
	}

	return json.Marshal(cape)


}

func SetType(cape *models.Cape) error {

	_, err := db.Db.Exec("UPDATE capes SET type=? WHERE uuid=?", cape.CapeType, cape.Uuid)
	if err != nil {
		log.Println(err)
	}

	return err
}

func GetCapesForm() []models.Cape {
	var (
		cape  models.Cape
		capes []models.Cape
	)

	rows, err := db.Db.Query("SELECT * FROM capes")

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		rows.Scan(&cape.Uuid, &cape.CapeType)
		capes = append(capes, cape)
	}

	defer rows.Close()
	return capes
}