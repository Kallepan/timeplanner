package mock

import "github.com/gin-gonic/gin"

type SystemControllerMock struct {
}

func (m *SystemControllerMock) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong"})
}
