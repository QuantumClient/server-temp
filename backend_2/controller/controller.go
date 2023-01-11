package controller

import "github.com/gookit/rux"

type WebController interface {
	AddRoutes(router *rux.Router)
}
