package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemControllerMock struct {
}

func (m *SystemControllerMock) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}
