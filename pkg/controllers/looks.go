package controllers

import (
	"look-api/pkg/models"

	"github.com/juansgt/generics/services"
	"github.com/juansgt/model-test/v2/dataAccess"
)

type LooksController struct {
	findLooksQueryService services.IQueryServiceNoInput[[]dataAccess.Look]
}

func NewLooksController(findLooksQueryService services.IQueryServiceNoInput[[]dataAccess.Look]) *LooksController {
	return &LooksController{
		findLooksQueryService: findLooksQueryService,
	}
}

func (looksController *LooksController) GetLooks() []models.Look {
	var looks []dataAccess.Look = make([]dataAccess.Look, 0)
	var apiLooks []models.Look = make([]models.Look, 0)

	looks = looksController.findLooksQueryService.Execute()

	for _, look := range looks {
		apiLooks = append(apiLooks, models.NewLook(look.Id(), look.Name, look.Brand))
	}

	return apiLooks
}
