package mock

import "github.com/gin-gonic/gin"

type PersonRelControllerMock struct {
}

func (m *PersonRelControllerMock) AddAbsency(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "AddAbsency"})
}

func (m *PersonRelControllerMock) RemoveAbsency(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "RemoveAbsency"})
}

func (m *PersonRelControllerMock) FindAbsencyForPerson(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "FindAbsencyForPerson"})
}

func (m *PersonRelControllerMock) AddDepartment(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "AddDepartment"})
}

func (m *PersonRelControllerMock) RemoveDepartment(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "RemoveDepartment"})
}

func (m *PersonRelControllerMock) AddWorkplace(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "AddWorkplace"})
}

func (m *PersonRelControllerMock) RemoveWorkplace(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "RemoveWorkplace"})
}

func (m *PersonRelControllerMock) AddWeekday(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "AddWeekday"})
}

func (m *PersonRelControllerMock) RemoveWeekday(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "RemoveWeekday"})
}
