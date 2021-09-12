package models

import (
	"backend/db"
	"github.com/google/uuid"
	"log"
)

type Permission struct {
	ID        uuid.UUID  `json:"uuid"`
	Username  string     `json:"username"`
	Admin     bool       `json:"admin"`
	Access    bool       `json:"access"`
	Hwid      NullString `json:"hwid"`
	CreatedAt string     `json:"created_at"` //time was giving dumb errors with mysql
}

func PermsfromUser(user *User) *Permission {
	perms := &Permission{}

	perms.ID = user.Uuid
	perms.Username = user.Username

	res, err := db.Db.Query("SELECT * FROM permissions WHERE uuid=?", perms.ID)
	defer res.Close()
	if err != nil {
		log.Println(err)
	}

	if res.Next() {
		err := res.Scan(&perms.ID, &perms.Admin, &perms.Access, &perms.Hwid, &perms.CreatedAt)
		if err != nil {
			log.Println(err)
		}
	} else {
		st, err := db.Db.Prepare("INSERT into permissions SET uuid=?")
		if err != nil {
			log.Println(err)
		}

		_, err = st.Exec(perms.ID)

		if err != nil {
			log.Println(err)
		}
	}
	return perms
}
