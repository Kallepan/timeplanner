package mock

import "github.com/gin-gonic/gin"

type WeekdayControllerMock struct {
}

func (m *WeekdayControllerMock) AddWeekdayToTimeslot(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "AddWeekdayToTimeslot"})
}

func (m *WeekdayControllerMock) RemoveWeekdayFromTimeslot(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "RemoveWeekdayFromTimeslot"})
}
