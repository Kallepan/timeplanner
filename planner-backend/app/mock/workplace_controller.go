package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorkplaceControllerMock struct {
}

func (m *WorkplaceControllerMock) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "GetAll"})
}

func (m *WorkplaceControllerMock) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get"})
}

func (m *WorkplaceControllerMock) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Create"})
}

func (m *WorkplaceControllerMock) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update"})
}

func (m *WorkplaceControllerMock) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete"})
}
