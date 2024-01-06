package repository

import (
	"auth-backend/app/domain/dao"
	"log/slog"

	"github.com/google/uuid"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	FindAllPermissions() ([]dao.Permission, error)
	FindPermissionById(id uuid.UUID) (dao.Permission, error)
	Save(Permission *dao.Permission) (dao.Permission, error)
	DeletePermissionById(id uuid.UUID) error
}

type PermissionRepositoryImpl struct {
	db *gorm.DB
}

func (r PermissionRepositoryImpl) FindAllPermissions() ([]dao.Permission, error) {
	var Permissions []dao.Permission

	var err = r.db.Find(&Permissions).Error
	if err != nil {
		slog.Error("Got an error finding all couples.", "error", err)
		return nil, err
	}

	return Permissions, nil
}

func (r PermissionRepositoryImpl) FindPermissionById(id uuid.UUID) (dao.Permission, error) {
	Permission := dao.Permission{
		BaseModel: dao.BaseModel{
			ID: id,
		},
	}
	err := r.db.First(&Permission).Error
	if err != nil {
		slog.Error("Got and error when find Permission by id.", "error", err)
		return dao.Permission{}, err
	}
	return Permission, nil
}

func (r PermissionRepositoryImpl) Save(Permission *dao.Permission) (dao.Permission, error) {
	var err = r.db.Save(Permission).Error
	if err != nil {
		slog.Error("Got an error when save Permission.", "error", err)
		return dao.Permission{}, err
	}
	return *Permission, nil
}

func (r PermissionRepositoryImpl) DeletePermissionById(id uuid.UUID) error {
	err := r.db.Delete(&dao.Permission{}, id).Error
	if err != nil {
		slog.Error("Got an error when delete Permission.", "error", err)
		return err
	}
	return nil
}

func PermissionRepositoryInit(db *gorm.DB) *PermissionRepositoryImpl {
	db.AutoMigrate(&dao.Permission{})
	return &PermissionRepositoryImpl{
		db: db,
	}
}

var permissionRepositorySet = wire.NewSet(
	PermissionRepositoryInit,
	wire.Bind(new(PermissionRepository), new(*PermissionRepositoryImpl)),
)
