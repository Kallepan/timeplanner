package controller

import (
	"auth-backend/app/service"

	"github.com/gin-gonic/gin"
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
	svc service.UserService
}

func (u UserControllerImpl) GetAll(ctx *gin.Context) {
	u.svc.GetAllUsers(ctx)
}

func (u UserControllerImpl) Get(ctx *gin.Context) {
	u.svc.GetUserById(ctx)
}

func (u UserControllerImpl) Create(ctx *gin.Context) {
	u.svc.AddUser(ctx)
}

func (u UserControllerImpl) Update(ctx *gin.Context) {
	u.svc.UpdateUser(ctx)
}

func (u UserControllerImpl) Delete(ctx *gin.Context) {
	u.svc.DeleteUser(ctx)
}

func (u UserControllerImpl) AddPermission(ctx *gin.Context) {
	u.svc.AddPermission(ctx)
}

func (u UserControllerImpl) DeletePermission(ctx *gin.Context) {
	u.svc.DeletePermission(ctx)
}

func UserControllerInit(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: userService,
	}
}