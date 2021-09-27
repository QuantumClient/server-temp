package controllers

import (
	"backend/models"
	"time"
)

var limmit time.Duration = 6 * time.Hour
var onlines []*models.Online

func GetAllOnline() []*models.Online {
	for i := range onlines {
		var o = onlines[i]
		if time.Now().UnixNano() > o.Expiration {
			RemoveFromOnline(o)
		}
	}
	return onlines
}

func AddToOnline(o *models.Online) []*models.Online {
	index := findIndex(o)
	o.Expiration = time.Now().Add(limmit).UnixNano()
	if index == -1 {
		onlines = append(onlines, o)
	} else {
		onlines[index].Expiration = o.Expiration
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
