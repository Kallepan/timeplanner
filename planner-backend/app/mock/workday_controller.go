package mock

import "github.com/gin-gonic/gin"

type WorkdayControllerMock struct {
}

func (m *WorkdayControllerMock) GetWorkdaysForDepartmentAndDate(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetWorkdaysForDepartmentAndDate"})
}

func (m *WorkdayControllerMock) GetWorkday(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetWorkday"})
}

func (m *WorkdayControllerMock) AssignPersonToWorkday(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "AssignPersonToWorkday"})
}

func (m *WorkdayControllerMock) UnassignPersonFromWorkday(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "UnassignPersonFromWorkday"})
}
