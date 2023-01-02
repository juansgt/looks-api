package main

import (
	"look-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juansgt/generics/services"
	"github.com/juansgt/model-test/v2/dataAccess"
	"github.com/juansgt/model-test/v2/services/findLooksService"
)

func getLooks(context *gin.Context) {
	var findLooksQueryService services.IQueryServiceNoInput[[]dataAccess.Look]
	var looks []dataAccess.Look = make([]dataAccess.Look, 0)
	var apiLooks []models.Look = make([]models.Look, 0)

	findLooksQueryService = findLooksService.NewFindLooksQueryService(new(dataAccess.LookRepositoryMock))
	looks = findLooksQueryService.Execute()

	for _, look := range looks {
		apiLooks = append(apiLooks, models.NewLook(look.Id(), look.Name, look.Brand))
	}

	context.IndentedJSON(http.StatusOK, apiLooks)
}

func main() {
	ginEngine := gin.Default()
	ginEngine.GET("/looks", getLooks)
	ginEngine.Run("localhost:9090")
}
