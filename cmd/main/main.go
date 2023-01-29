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

var looksController *controllers.LooksController
var database *mongo.Database
var mongoClient *mongo.Client

func initializeDependencies() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://atlasAdmin:Cripto0Virtual@cluster0.yolpv.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	mongoClient = client
	database = mongoClient.Database("wap")
	looksController = controllers.NewLooksController(findLooksService.NewFindLooksQueryService(lookRepository.NewLookRepositoryMongodb(database)))
}

func main() {
	initializeDependencies()
	ginEngine := gin.Default()
	ginEngine.GET("/looks", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, looksController.GetLooks())
	})
	ginEngine.Run("localhost:9090")
}
