package controllers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"log"
)

func GetProjects() []models.Project {
	var (
		project models.Project
		projects []models.Project
	)

	rows, err := db.Db.Query("SELECT * FROM projects")

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		rows.Scan(&project.Name ,&project.Verison, &project.GitHub, &project.Link)
		projects = append(projects, project)
	}
	defer rows.Close()

	return projects

}

func UpdateProjectVersion(project *models.Project) error {

	_, err := db.Db.Exec("UPDATE projects SET version=? WHERE name=?", project.Verison, project.Name)
	if err != nil {
		log.Println(err)
	}

	return err
}

func GetProject(project *models.Project) ([]byte, error) {
	res, err := db.Db.Query("SELECT * FROM projects WHERE name=?", project.Name)
	defer res.Close()
	if err != nil {
		log.Println(err)
	}

	if res.Next() {
		err := res.Scan(&project.Name, &project.Verison, &project.GitHub, &project.Link)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		return nil, nil
	}

	return json.Marshal(project)

}

func UpdateProjectLink(project *models.Project) error {

	_, err := db.Db.Exec("UPDATE projects SET link=? WHERE name=?", project.Link, project.Name)
	if err != nil {
		log.Println(err)
	}

	return err
}