/**
* This package handles the authentication logic of the application.
* It has three main functions:
* - Login: Generate a JWT token
* - Me: Get the user data from the JWT token
* - Logout: invalidate a JWT token
**/
package service

import (
	"api-gateway/app/constant"
	"api-gateway/app/domain/dao"
	"api-gateway/app/domain/dco"
	"api-gateway/app/pkg"
	"api-gateway/app/repository"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	/**
	* Take in the username and password from the request body
	* Check if the username and password are correct
	* If correct, generate a JWT token and return it to the client via httpOnly cookie
	**/

	defer pkg.PanicHandler(c)
	slog.Info("start to execute program login")

	var request dco.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Error happened: when mapping request", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	user, err := a.UserRepository.FindUserByUsername(request.Username)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		slog.Error("Error happened: when get data from database", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error happened: when get data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		slog.Error("Error happened: when compare password", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dco.JWTClaim{
		Username:   user.Username,
		Department: user.Department.ID.String(),
		IsAdmin:    user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(dco.JWTExpirationTime)),
		},
	})

	// sign the token with the secret key
	tokenString, err := token.SignedString(dco.JWTSigningKey)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	// set the token in the cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, int(dco.JWTExpirationTime.Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, mapUserToUserResponse(user)))
}

func (a AuthServiceImpl) Me(c *gin.Context) {
	/**
	* This function gets the Cookie from the request header and checks it for validity.
	* If the token is valid, it returns the user data to be used in the frontend.
	**/
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program me")

	// get the token from the cookie
	claim, exists := c.Get("retrievedToken")
	if !exists {
		slog.Error("Error happened: when get token from cookie", "error", "token not found")
		pkg.PanicException(constant.Unauthorized)
	}

	// find the user data from the database
	data, err := a.UserRepository.FindUserByUsername(claim.(*dco.JWTClaim).Username)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		slog.Error("Error happened: when get data from database", "error", err)
		pkg.PanicException(constant.Unauthorized)
	default:
		slog.Error("Error happened: when get data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	// if departmentQuery parameter is not "", check if user belongs to the department
	departmentQuery := c.Query("department")
	if departmentQuery != "" && departmentQuery != data.Department.Name && !data.IsAdmin {
		pkg.PanicException(constant.Unauthorized)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, mapUserToAuthResponse(data)))
}

func (a AuthServiceImpl) Logout(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program logout")

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.Status(http.StatusOK)
}

var authServiceSet = wire.NewSet(
	wire.Struct(new(AuthServiceImpl), "*"),
	wire.Bind(new(AuthService), new(*AuthServiceImpl)),
)

func mapUserToAuthResponse(user dao.User) dco.AuthResponse {
	/**
	* This function maps the user data from the database to the response data
	**/
	return dco.AuthResponse{
		Username: user.Username,
		Email:    user.Email,
	}
}
