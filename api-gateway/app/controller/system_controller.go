package controller

import (
	"api-gateway/app/constant"
	"api-gateway/app/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type SystemController interface {
	Ping(ctx *gin.Context)
}

type SystemControllerImpl struct{}

func (s SystemControllerImpl) Ping(c *gin.Context) {
	defer pkg.PanicHandler(c)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

var systemControllerSet = wire.NewSet(
	wire.Struct(new(SystemControllerImpl), "*"),
	wire.Bind(new(SystemController), new(*SystemControllerImpl)),
)
