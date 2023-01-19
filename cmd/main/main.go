package main

import (
	"look-api/pkg/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juansgt/model-test/v3/dataAccess/lookRepository"
	"github.com/juansgt/model-test/v3/services/findLooksService"
)

var looksController *controllers.LooksController

func initializeDependencies() {
	looksController = controllers.NewLooksController(findLooksService.NewFindLooksQueryService(new(lookRepository.LookRepositoryMock)))
}

func main() {
	initializeDependencies()
	ginEngine := gin.Default()
	ginEngine.GET("/looks", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, looksController.GetLooks())
	})
	ginEngine.Run("localhost:9090")
}
