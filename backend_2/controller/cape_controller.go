package controller

import (
	"github.com/google/uuid"
	"github.com/gookit/rux"
	"quantumclient.org/backend/v2/models"
	"quantumclient.org/backend/v2/services"
)

type CapeController struct {
	capeService *services.CapeService
}

// NewCapeController returns a new instance of CapeController.
func NewCapeController(capeService *services.CapeService) *CapeController {
	return &CapeController{capeService: capeService}
}

func (c CapeController) AddRoutes(r *rux.Router) {
	r.GET("/", c.getEnabledCapes)
	r.GET("/all", c.getCapes)
	r.GET("/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", c.getCape)
	r.PUT("/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", c.updateCape)
	r.POST("/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}/enabled", c.enableCape)
}

func (c CapeController) getCape(ctx *rux.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	cape, err := c.capeService.GetCape(id)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	ctx.JSON(200, cape)
}

func (c CapeController) getCapes(ctx *rux.Context) {
	capes, err := c.capeService.GetCapes()
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	ctx.JSON(200, capes)
}

func (c CapeController) getEnabledCapes(ctx *rux.Context) {
	capes, err := c.capeService.GetEnabledCapes()
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	ctx.JSON(200, capes)
}

func (c CapeController) updateCape(ctx *rux.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	var cape models.Cape
	if err = ctx.Bind(&cape); err != nil {
		ctx.JSON(400, err)
		return
	}
	cape.Uuid = id
	user, has := ctx.Get("user")
	if !has {
		ctx.JSON(400, "user not found")
		return
	}
	err = c.capeService.UpdateCape(user.(*models.User).Uuid, cape)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	ctx.JSON(200, cape)
}

func (c CapeController) enableCape(ctx *rux.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	user, has := ctx.Get("user")
	if !has || user.(*models.User).Uuid != id {
		ctx.JSON(400, "user not found")
		return
	}
	err = c.capeService.SetCapeEnabled(id)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	ctx.JSON(200, id)
}