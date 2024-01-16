package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DepartmentControllerMock struct {
}

func (m *DepartmentControllerMock) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "GetAll"})
}

func (m *DepartmentControllerMock) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get"})
}

func (m *DepartmentControllerMock) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Create"})
}

func (m *DepartmentControllerMock) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update"})
}

func (m *DepartmentControllerMock) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete"})
}
