package service

import (
	"auth-backend/app/constant"
	"auth-backend/app/domain/dao"
	"auth-backend/app/pkg"
	"auth-backend/app/repository"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/google/wire"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	AddPermission(c *gin.Context)
	DeletePermission(c *gin.Context)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (u UserServiceImpl) UpdateUser(c *gin.Context) {
	/* Method to update user data by id */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program update user data by id")

	id := c.Param("userID")
	userID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Error happened: when mapping request", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.UserRepository.FindUserById(userID)
	if err != nil {
		slog.Error("Error happened: when get data from database", "error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	// Foreign keys
	data.DepartmentID = request.DepartmentID

	// Data
	data.Email = request.Email
	data.Username = request.Username

	// Save to database
	data, err = u.UserRepository.Save(&data)
	if err != nil {
		slog.Error("Error happened: when updating data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetUserById(c *gin.Context) {
	/* Method to get user data by id */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get user by id")

	id := c.Param("userID")
	userID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error happened: when parsing uuid", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.UserRepository.FindUserById(userID)
	if err != nil {
		slog.Error("Error happened: when get data from database", "error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) AddUser(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add data user")

	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Error happened: when mapping request", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)

	data, err := u.UserRepository.Save(&request)
	if err != nil {
		slog.Error("Error happened: when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetAllUsers(c *gin.Context) {
	/* Method to get all user data */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute get all data user")

	data, err := u.UserRepository.FindAllUsers()
	if err != nil {
		slog.Error("Error happened: when find all user data", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) DeleteUser(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute delete data user by id")

	id := c.Param("userID")
	userID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error happened: when parsing string to int", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	err = u.UserRepository.DeleteUser(userID)
	if err != nil {
		slog.Error("Error happened: when try delete data user from DB", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func (u UserServiceImpl) AddPermission(c *gin.Context) {
	/* Method to add permission to user */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add permission to user")

	id := c.Param("userID")
	userID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error happened: when parsing uuid", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	pId := c.Param("permissionID")
	permissionID, err := uuid.Parse(pId)
	if err != nil {
		slog.Error("Error happened: when parsing uuid", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	err = u.UserRepository.AddPermissionToUser(userID, permissionID)
	if err != nil {
		slog.Error("Error happened: when add permission to user", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func (u UserServiceImpl) DeletePermission(c *gin.Context) {
	/* Method to delete permission to user */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete permission to user")

	id := c.Param("userID")
	userID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error happened: when parsing uuid", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	pId := c.Param("permissionID")
	permissionID, err := uuid.Parse(pId)
	if err != nil {
		slog.Error("Error happened: when parsing uuid", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	err = u.UserRepository.DeletePermissionFromUser(userID, permissionID)
	if err != nil {
		slog.Error("Error happened: when delete permission to user", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

var userServiceSet = wire.NewSet(
	wire.Struct(new(UserServiceImpl), "*"),
	wire.Bind(new(UserService), new(*UserServiceImpl)),
)
