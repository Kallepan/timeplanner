package controller

import (
	"auth-backend/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type UserController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)

	AddPermission(ctx *gin.Context)
	DeletePermission(ctx *gin.Context)
}

type UserControllerImpl struct {
	UserService service.UserService
}

func (u UserControllerImpl) GetAll(ctx *gin.Context) {
	u.UserService.GetAllUsers(ctx)
}

func (u UserControllerImpl) Get(ctx *gin.Context) {
	u.UserService.GetUserById(ctx)
}

func (u UserControllerImpl) Create(ctx *gin.Context) {
	u.UserService.AddUser(ctx)
}

func (u UserControllerImpl) Update(ctx *gin.Context) {
	u.UserService.UpdateUser(ctx)
}

func (u UserControllerImpl) Delete(ctx *gin.Context) {
	u.UserService.DeleteUser(ctx)
}

func (u UserControllerImpl) AddPermission(ctx *gin.Context) {
	u.UserService.AddPermission(ctx)
}

func (u UserControllerImpl) DeletePermission(ctx *gin.Context) {
	u.UserService.DeletePermission(ctx)
}

var userControllerSet = wire.NewSet(
	wire.Struct(new(UserControllerImpl), "*"),
	wire.Bind(new(UserController), new(*UserControllerImpl)),
)