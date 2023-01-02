package main

import (
	"look-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juansgt/model-test/dataAccess"
	"github.com/juansgt/model-test/services"
)

func getLooks(context *gin.Context) {
	var findLooksQueryService *services.FindLooksQueryService
	var looks []dataAccess.Look = make([]dataAccess.Look, 0)
	var apiLooks []models.Look = make([]models.Look, 0)

	findLooksQueryService = services.NewFindLooksQueryService(new(dataAccess.LookRepositoryMock))
	looks = findLooksQueryService.FindLooks()

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
