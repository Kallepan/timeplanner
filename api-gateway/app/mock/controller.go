package mock

import "github.com/gin-gonic/gin"

/** SystemControllerMock */
type SystemControllerMock struct {
}

func (m *SystemControllerMock) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong"})
}

/** UserControllerMock */
type UserControllerMock struct{}

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

func (m *UserControllerMock) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Login"})
}

func (m *UserControllerMock) Me(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Me"})
}

func (m *UserControllerMock) Logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Logout"})
}

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

/** PermissionControllerMock */
type PermissionControllerMock struct{}

func (m *PermissionControllerMock) GetAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetAll"})
}

func (m *PermissionControllerMock) Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Get"})
}

func (m *PermissionControllerMock) Create(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Create"})
}

func (m *PermissionControllerMock) Update(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Update"})
}

func (m *PermissionControllerMock) Delete(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Delete"})
}
