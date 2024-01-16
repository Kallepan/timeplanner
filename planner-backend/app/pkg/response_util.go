package pkg

import (
	"planner-backend/app/constant"
	"planner-backend/app/domain/dto"
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
