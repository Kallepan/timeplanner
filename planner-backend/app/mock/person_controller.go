package mock

import "github.com/gin-gonic/gin"

type PersonControllerMock struct {
}

func (m *PersonControllerMock) GetAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetAll"})
}

func (m *PersonControllerMock) Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Get"})
}

func (m *PersonControllerMock) Create(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Create"})
}

func (m *PersonControllerMock) Update(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Update"})
}

func (m *PersonControllerMock) Delete(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Delete"})
}
