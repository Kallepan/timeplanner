package service

import (
	"auth-backend/app/constant"
	"auth-backend/app/domain/dao"
	"auth-backend/app/pkg"
	"auth-backend/app/repository"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/google/wire"
)

type PermissionService interface {
	GetAllPermissions(c *gin.Context)
	GetPermissionById(c *gin.Context)
	AddPermission(c *gin.Context)
	UpdatePermission(c *gin.Context)
	DeletePermission(c *gin.Context)
}

type PermissionServiceImpl struct {
	PermissionRepository repository.PermissionRepository
}

func (p PermissionServiceImpl) GetAllPermissions(c *gin.Context) {
	/* GetAllPermissions is a function to get all permissions
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all permissions")

	data, err := p.PermissionRepository.FindAllPermissions()
	if err != nil {
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) GetPermissionById(c *gin.Context) {
	/* GetPermissionById is a function to get permission by id
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get permission by id")

	id := c.Param("permissionID")
	permissionID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := p.PermissionRepository.FindPermissionById(permissionID)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) AddPermission(c *gin.Context) {
	/* AddPermission is a function to add permission
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add permission")

	var request dao.Permission
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Error when binding json", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := p.PermissionRepository.Save(&request)
	if err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) UpdatePermission(c *gin.Context) {
	/* UpdatePermission is a function to update permission by id
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program update permission")

	id := c.Param("permissionID")
	permissionID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	var request dao.Permission
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Error when binding json", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := p.PermissionRepository.FindPermissionById(permissionID)
	if err != nil {
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.Name = request.Name
	data.Description = request.Description

	data, err = p.PermissionRepository.Save(&data)
	if err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) DeletePermission(c *gin.Context) {
	/* DeletePermission is a function to delete permission by id
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete permission")

	id := c.Param("permissionID")
	permissionID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	err = p.PermissionRepository.DeletePermissionById(permissionID)
	if err != nil {
		slog.Error("Error when deleting data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

var permissionServiceSet = wire.NewSet(
	wire.Struct(new(PermissionServiceImpl), "*"),
	wire.Bind(new(PermissionService), new(*PermissionServiceImpl)),
)
