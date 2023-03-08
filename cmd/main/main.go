package main

import (
	"context"
	"look-api/pkg/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juansgt/model-test/v3/dataAccess/lookRepository"
	"github.com/juansgt/model-test/v3/services/findLooksService"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func getLooksController() *controllers.LooksController {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://thewappcontact:MariaJuanPaula@the-wap.oriorrs.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	database = client.Database("wap")

	return controllers.NewLooksController(findLooksService.NewFindLooksQueryService(lookRepository.NewLookRepositoryMongodb(database)))
}

func main() {
	ginEngine := gin.Default()
	ginEngine.GET("/looks", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, getLooksController().GetLooks())
	})
	ginEngine.Run()
}
