package controllers

import (
	"look-api/pkg/models"

	"github.com/juansgt/generics/services"
	"github.com/juansgt/looks-model/v3/dataAccess/lookRepository"
)

type LooksController struct {
	findLooksQueryService services.IQueryServiceNoInput[[]lookRepository.Look]
}

func NewLooksController(findLooksQueryService services.IQueryServiceNoInput[[]lookRepository.Look]) *LooksController {
	return &LooksController{
		findLooksQueryService: findLooksQueryService,
	}
}

func (looksController *LooksController) GetLooks() []models.Look {
	var looks []lookRepository.Look = make([]lookRepository.Look, 0)
	var apiLooks []models.Look = make([]models.Look, 0)

	looks = looksController.findLooksQueryService.Execute()

	for _, look := range looks {
		apiLooks = append(apiLooks, models.NewLook(look.Id(), look.Colour, look.Brand))
	}

	return apiLooks
}
