package mock

import "github.com/gin-gonic/gin"

type UserControllerMock struct{}

/* Controller interface implementations */
func (m *UserControllerMock) GetAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetAll"})
}

func (m *UserControllerMock) Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Get"})
}

func (m *UserControllerMock) Create(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Create"})
}

func (m *UserControllerMock) Update(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Update"})
}

func (m *UserControllerMock) Delete(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Delete"})
}

func (m *UserControllerMock) AddPermission(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "AddPermission"})
}

func (m *UserControllerMock) DeletePermission(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "DeletePermission"})
}
