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

	slog.Info("start to compare password", "hashedPassword", user.Password, "password", request.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		slog.Error("Error happened: when compare password", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	// generate JWT token
	convertedPermissions := make([]string, len(user.Permissions))
	for i, v := range user.Permissions {
		convertedPermissions[i] = v.ConvertToID()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dco.JWTClaim{
		Username:    user.Username,
		Department:  user.Department.ID.String(),
		Permissions: convertedPermissions,
		IsAdmin:     user.IsAdmin,
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
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, gin.H{
		"message": "login successfully",
	}))
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
	rawData, err := a.UserRepository.FindUserByUsername(claim.(*dco.JWTClaim).Username)
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

	data := mapUserToUserResponse(rawData)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (a AuthServiceImpl) Logout(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program logout")

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, gin.H{
		"message": "logout successfully",
	}))
}

var authServiceSet = wire.NewSet(
	wire.Struct(new(AuthServiceImpl), "*"),
	wire.Bind(new(AuthService), new(*AuthServiceImpl)),
)
