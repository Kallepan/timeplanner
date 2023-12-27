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
)

type PermissionService interface {
	GetAllPermissions(c *gin.Context)
	GetPermissionById(c *gin.Context)
	AddPermission(c *gin.Context)
	UpdatePermission(c *gin.Context)
	DeletePermission(c *gin.Context)
}

type PermissionServiceImpl struct {
	permissionRepository repository.PermissionRepository
}

func (p PermissionServiceImpl) GetAllPermissions(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all permissions")

	data, err := p.permissionRepository.FindAllPermissions()
	if err != nil {
		slog.Error("Error when fetching data from database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) GetPermissionById(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get permission by id")

	id := c.Param("permissionID")
	permissionID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := p.permissionRepository.FindPermissionById(permissionID)
	if err != nil {
		slog.Error("Error when fetching data from database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) AddPermission(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add permission")

	var request dao.Permission
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Error when binding json", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := p.permissionRepository.Save(&request)
	if err != nil {
		slog.Error("Error when saving data to database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) UpdatePermission(c *gin.Context) {
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

	data, err := p.permissionRepository.FindPermissionById(permissionID)
	if err != nil {
		slog.Error("Error when fetching data from database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	data.Name = request.Name
	data.Description = request.Description

	data, err = p.permissionRepository.Save(&data)
	if err != nil {
		slog.Error("Error when saving data to database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) DeletePermission(c *gin.Context) {
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete permission")

	id := c.Param("permissionID")
	permissionID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	err = p.permissionRepository.DeletePermissionById(permissionID)
	if err != nil {
		slog.Error("Error when deleting data from database", err)
		pkg.PanicException(constant.DatabaseError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func PermissionServiceInit(permissionRepository repository.PermissionRepository) *PermissionServiceImpl {
	return &PermissionServiceImpl{
		permissionRepository: permissionRepository,
	}
}