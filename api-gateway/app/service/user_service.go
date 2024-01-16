package service

import (
	"api-gateway/app/constant"
	"api-gateway/app/domain/dao"
	"api-gateway/app/domain/dco"
	"api-gateway/app/pkg"
	"api-gateway/app/repository"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/google/wire"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	var rawRequest dco.UserRequest
	if err := c.ShouldBindJSON(&rawRequest); err != nil {
		slog.Error("Error happened: when mapping request", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	request := mapUserRequestToUser(rawRequest)

	oldData, err := u.UserRepository.FindUserById(userID)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error happened: when get data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	// Foreign keys
	oldData.DepartmentID = request.DepartmentID

	// Data
	oldData.Email = request.Email
	oldData.Username = request.Username

	// Save to database
	rawData, err := u.UserRepository.Save(&oldData)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error happened: when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}
	data := mapUserToUserResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetUserById(c *gin.Context) {
	/* Method to get user data by id */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get user by id")

	id := c.Param("userID")
	userID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error happened when parsing uuid", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := u.UserRepository.FindUserById(userID)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error happened: when get data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}
	data := mapUserToUserResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) AddUser(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add data user")

	var rawRequest dco.UserRequest
	if err := c.ShouldBindJSON(&rawRequest); err != nil {
		slog.Error("Error happened: when mapping request", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	request := mapUserRequestToUser(rawRequest)

	// Check if username already exist
	_, err := u.UserRepository.FindUserByUsername(request.Username)
	switch err {
	case nil:
		pkg.PanicException(constant.Conflict)
	case gorm.ErrRecordNotFound:
		break
	default:
		slog.Error("Error happened: when get data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	// Hash password
	if hash, err := bcrypt.GenerateFromPassword([]byte(rawRequest.Password), 15); err != nil {
		slog.Error("Error happened: when hashing password", "error", err)
		pkg.PanicException(constant.UnknownError)
	} else {
		request.Password = string(hash)
	}

	rawData, err := u.UserRepository.Save(&request)
	switch err {
	case nil:
		break
	default:
		slog.Error("Error happened: when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}
	data := mapUserToUserResponse(rawData)

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetAllUsers(c *gin.Context) {
	/* Method to get all user data */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute get all data user")

	rawData, err := u.UserRepository.FindAllUsers()
	if err != nil {
		slog.Error("Error happened: when find all user data", "error", err)
		pkg.PanicException(constant.UnknownError)
	}
	data := mapUserListToUserResponseList(rawData)

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
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error happened: when delete data user from DB", "error", err)
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
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error happened: when add permission to user", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, pkg.Null()))
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
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		pkg.PanicException(constant.InvalidRequest)
	default:
		slog.Error("Error happened: when delete permission to user", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

var userServiceSet = wire.NewSet(
	wire.Struct(new(UserServiceImpl), "*"),
	wire.Bind(new(UserService), new(*UserServiceImpl)),
)

func mapUserToUserResponse(user dao.User) dco.UserResponse {
	/* mapUserToUserResponse is a function to map user to user response
	 * @param user is dao.User
	 * @return dco.UserResponse
	 */
	return dco.UserResponse{
		BaseModel: dco.BaseModel{
			ID:        user.BaseModel.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Username: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		Department: dco.DepartmentResponse{
			BaseModel: dco.BaseModel{
				ID:        user.Department.BaseModel.ID,
				CreatedAt: user.Department.CreatedAt,
				UpdatedAt: user.Department.UpdatedAt,
			},
			Name: user.Department.Name,
		},
	}
}

func mapUserListToUserResponseList(users []dao.User) []dco.UserResponse {
	/* mapUserListToUserResponseList is a function to map user list to user response list
	 * @param users is []dao.User
	 * @return []dco.UserResponse
	 */
	var result []dco.UserResponse
	for _, user := range users {
		result = append(result, mapUserToUserResponse(user))
	}
	return result
}

func mapUserRequestToUser(req dco.UserRequest) dao.User {
	/* mapUserRequestToUser is a function to map user request to user
	 * @param req is dco.UserRequest
	 * @return dao.User
	 */
	return dao.User{
		Username:     req.Username,
		Email:        req.Email,
		Password:     req.Password,
		DepartmentID: req.DepartmentID,
		IsAdmin:      req.IsAdmin,
	}
}
