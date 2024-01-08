package pkg

import (
	"api-gateway/app/constant"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constant.ResponseStatus) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}

func NewException(responseKey constant.ResponseStatus) error {
	err := errors.New(responseKey.GetResponseMessage())
	err = fmt.Errorf("%s: %w", responseKey.GetResponseStatus(), err)
	return err
}

func PanicHandler(ctx *gin.Context) {
	if err := recover(); err != nil {
		// get error message
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")
		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		// log error
		slog.Error("Error when executing program", "error", err)

		// handle error
		switch key {
		case
			constant.DataNotFound.GetResponseStatus():
			ctx.JSON(http.StatusNotFound, BuildResponse_(key, msg, Null()))
			ctx.Abort()
		case
			constant.Unauthorized.GetResponseStatus():
			ctx.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, Null()))
			ctx.Abort()
		case
			constant.InvalidRequest.GetResponseStatus():
			ctx.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
			ctx.Abort()
		default:
			ctx.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, Null()))
			ctx.Abort()
		}
	}
}
