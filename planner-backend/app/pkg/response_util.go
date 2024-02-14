package pkg

import (
	"net/http"
	"planner-backend/app/constant"
	"planner-backend/app/domain/dto"

	"github.com/gin-gonic/gin"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.APIResponse[T] {
	/** This function is used to build response */
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponseWithCustomMessage[T any](responseStatus constant.ResponseStatus, message string, data T) {
	/** This function is used to build response with custom message */
	BuildResponse_(responseStatus.GetResponseStatus(), message, data)
}

func BuildResponse_[T any](status string, message string, data T) dto.APIResponse[T] {
	/** This function is used to build response */
	return dto.APIResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}

func SendResponse[T any](ctx *gin.Context, responseStatus constant.ResponseStatus, data T) {
	/** This function is used to send response */

	switch responseStatus.GetResponseStatus() {
	case constant.InvalidRequest.GetResponseStatus():
		ctx.JSON(http.StatusBadRequest, BuildResponse(responseStatus, data))
	default:
		ctx.JSON(http.StatusInternalServerError, BuildResponse(responseStatus, data))
	}
}
