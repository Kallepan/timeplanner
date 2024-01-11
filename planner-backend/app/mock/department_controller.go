package mock

import "github.com/gin-gonic/gin"

type DepartmentControllerMock struct {
}

func (m *DepartmentControllerMock) GetAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetAll"})
}

func (m *DepartmentControllerMock) Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Get"})
}

func (m *DepartmentControllerMock) Create(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Create"})
}

func (m *DepartmentControllerMock) Update(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Update"})
}

func (m *DepartmentControllerMock) Delete(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Delete"})
}
