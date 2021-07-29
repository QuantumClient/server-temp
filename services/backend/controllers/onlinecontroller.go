package controllers

import "backend/models"

var onlines []*models.Online

func GetAllOnline() []*models.Online {
	return onlines
}

func AddToOnline(o *models.Online) []*models.Online {
	index := findIndex(o)
	if index == -1 {
		onlines = append(onlines, o)
	}

	return onlines

}

func RemoveFromOnline(o *models.Online) []*models.Online {
	index := findIndex(o)
	if index != -1 {
		onlines = append(onlines[:index], onlines[index+1:]...)
	}
	return onlines
}

func findIndex(o *models.Online) int {
	for i, n := range onlines {
		if n.Uuid == o.Uuid {
			return i
		}
	}

	return -1
}
