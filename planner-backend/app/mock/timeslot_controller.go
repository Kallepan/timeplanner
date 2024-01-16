package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TimeslotControllerMock struct {
}

func (m *TimeslotControllerMock) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "GetAll"})
}

func (m *TimeslotControllerMock) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get"})
}

func (m *TimeslotControllerMock) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Create"})
}

func (m *TimeslotControllerMock) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update"})
}

func (m *TimeslotControllerMock) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete"})
}
