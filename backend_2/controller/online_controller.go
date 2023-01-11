package controller

import (
	"github.com/gookit/rux"
	"net/http"
	"quantumclient.org/backend/v2/models"
	"quantumclient.org/backend/v2/services"
)

type OnlineController struct {
	services.OnlineServiceInterface
}

func NewOnlineController() *OnlineController {
	return &OnlineController{
		OnlineServiceInterface: services.NewOnlineService(),
	}
}

func (c OnlineController) AddRoutes(r *rux.Router)  {
	r.GET("/", c.getOnlineAccounts)
	r.POST("/", c.addOnlineAccount)

}

func (c OnlineController) getOnlineAccounts(ctx *rux.Context)  {
	ctx.JSON(http.StatusCreated, c.OnlineServiceInterface.GetOnlineAccounts())
}

func (c OnlineController) addOnlineAccount(ctx *rux.Context)  {
	var onlineAccount models.Online
	err := ctx.AutoBind(&onlineAccount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	onlineAccount.Populate()

	r, err := c.OnlineServiceInterface.Add(&onlineAccount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, r)
}