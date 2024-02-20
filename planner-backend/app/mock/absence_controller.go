package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AbsenceControllerMock struct {
}

func (m *AbsenceControllerMock) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "GetAll"})
}
