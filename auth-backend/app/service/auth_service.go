/**
* This package handles the authentication logic of the application.
* It has three main functions:
* - Login: Generate a JWT token
* - Me: Get the user data from the JWT token
* - Logout: invalidate a JWT token
**/
package service

import (
	"auth-backend/app/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type AuthService interface {
	Login(c *gin.Context)
	Me(c *gin.Context)
	Logout(c *gin.Context)
}

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
}

func (a AuthServiceImpl) Login(c *gin.Context) {
	panic("implement me")
}

func (a AuthServiceImpl) Me(c *gin.Context) {
	panic("implement me")
}

func (a AuthServiceImpl) Logout(c *gin.Context) {
	panic("implement me")
}

var authServiceSet = wire.NewSet(
	wire.Struct(new(AuthServiceImpl), "*"),
	wire.Bind(new(AuthService), new(*AuthServiceImpl)),
)
