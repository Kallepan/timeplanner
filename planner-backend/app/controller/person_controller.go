package controller

import (
	"planner-backend/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type PersonController interface {
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type PersonControllerImpl struct {
	PersonService service.PersonService
}

func (u PersonControllerImpl) GetAll(ctx *gin.Context) {
	u.PersonService.GetAllPersons(ctx)
}

func (u PersonControllerImpl) Get(ctx *gin.Context) {
	u.PersonService.GetPersonByID(ctx)
}

func (u PersonControllerImpl) Create(ctx *gin.Context) {
	u.PersonService.AddPerson(ctx)
}

func (u PersonControllerImpl) Update(ctx *gin.Context) {
	u.PersonService.UpdatePerson(ctx)
}

func (u PersonControllerImpl) Delete(ctx *gin.Context) {
	u.PersonService.DeletePerson(ctx)
}

var personControllerSet = wire.NewSet(
	wire.Struct(new(PersonControllerImpl), "*"),
	wire.Bind(new(PersonController), new(*PersonControllerImpl)),
)
