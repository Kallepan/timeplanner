package pkg

import (
	"auth-backend/app/constant"
	"auth-backend/app/domain/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.APIResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) dto.APIResponse[T] {
	return dto.APIResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}