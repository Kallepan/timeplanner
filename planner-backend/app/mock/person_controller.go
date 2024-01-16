package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonControllerMock struct {
}

func (m *PersonControllerMock) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "GetAll"})
}

func (m *PersonControllerMock) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get"})
}

func (m *PersonControllerMock) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Create"})
}

func (m *PersonControllerMock) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update"})
}

func (m *PersonControllerMock) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete"})
}
