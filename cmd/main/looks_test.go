package main_test

import (
	"encoding/json"
	mock_services "look-api/mocks/queryService"
	"look-api/pkg/controllers"
	"look-api/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/juansgt/looks-model/v3/dataAccess/lookRepository"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestFindLooks_CorrectRequest_ReturnOkStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	var looks []models.Look
	queryService := mock_services.NewMockIQueryServiceNoInput[[]lookRepository.Look](ctrl)
	queryService.EXPECT().Execute().Return(findLooks())
	var looksController *controllers.LooksController = controllers.NewLooksController(queryService)

	router := SetUpRouter()

	router.GET("/looks", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, looksController.GetLooks())
	})
	req, _ := http.NewRequest("GET", "/looks", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	json.Unmarshal(recorder.Body.Bytes(), &looks)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.NotEmpty(t, looks)
}

func findLooks() []lookRepository.Look {
	var looks []lookRepository.Look
	var look *lookRepository.Look = &lookRepository.Look{
		Colour: "Dress",
		Brand:  "Bash",
	}
	looks = make([]lookRepository.Look, 0, 3)

	looks = append(looks, *look)

	return looks
}
