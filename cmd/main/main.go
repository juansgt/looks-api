package main

import (
	"context"
	"look-api/pkg/controllers"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/juansgt/authentication-service/authenticationService"
	"github.com/juansgt/http-services/httpServices"
	"github.com/juansgt/looks-model/v3/dataAccess/lookRepository"
	"github.com/juansgt/looks-model/v3/services/findLooksService"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
)

var database *mongo.Database
var looksController *controllers.LooksController
var httpService httpServices.IHttpService
var authService authenticationService.IAuthenticationService

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://thewappcontact:MariaJuanPaula@the-wap.oriorrs.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	database = client.Database("wap")
	looksController = controllers.NewLooksController(findLooksService.NewFindLooksQueryService(lookRepository.NewLookRepositoryMongodb(database)))
	httpService = httpServices.NewHttpServiceBase()
}

func main() {
	ginEngine := gin.Default()
	ginEngine.GET("/looks", authHandler, func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, looksController.GetLooks())
	})
	ginEngine.Run()
}

func authHandler(context *gin.Context) {
	opt := option.WithCredentialsFile("firebase-credentials-key.json") // Replace with the path to your service account key JSON file
	app, err := firebase.NewApp(context, nil, opt)
	if err != nil {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unable to run Authorization infraestructure"})
	}
	accessToken, err := httpService.GetAccessToken(context.Request)
	authService = authenticationService.NewFirebaseAuthenticationService(context, app)

	if err != nil {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
	}

	if authService.IsValidToken(accessToken) {
		context.Next() // Call the next handler in the chain
	} else {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Access Token"})
	}
}
