package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeekdayControllerMock struct {
}

func (m *WeekdayControllerMock) AddWeekdayToTimeslot(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "AddWeekdayToTimeslot"})
}

func (m *WeekdayControllerMock) RemoveWeekdayFromTimeslot(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "RemoveWeekdayFromTimeslot"})
}

func (m *WeekdayControllerMock) BulkUpdateWeekdaysForTimeslot(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "BulkUpdateWeekdaysForTimeslot"})
}
