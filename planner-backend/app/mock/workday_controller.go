package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorkdayControllerMock struct {
}

func (m *WorkdayControllerMock) GetWorkdaysForDepartmentAndDate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "GetWorkdaysForDepartmentAndDate"})
}

func (m *WorkdayControllerMock) GetWorkday(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "GetWorkday"})
}

func (m *WorkdayControllerMock) UpdateWorkday(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "UpdateWorkday"})
}

func (m *WorkdayControllerMock) AssignPersonToWorkday(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "AssignPersonToWorkday"})
}

func (m *WorkdayControllerMock) UnassignPersonFromWorkday(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "UnassignPersonFromWorkday"})
}
