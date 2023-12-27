package repository

import (
	"auth-backend/app/domain/dao"
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUsers() ([]dao.User, error)
	FindUserById(id uuid.UUID) (dao.User, error)
	Save(user *dao.User) (dao.User, error)
	DeleteUserById(id uuid.UUID) error

	AddPermissionToUser(userID uuid.UUID, permissionID uuid.UUID) error
	DeletePermissionFromUser(userID uuid.UUID, permissionID uuid.UUID) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) FindAllUsers() ([]dao.User, error) {
	var users []dao.User

	var err = u.db.Preload("Permission").Find(&users).Error
	if err != nil {
		slog.Error("Got an error finding all couples.", "error", err)
		return nil, err
	}

	return users, nil
}

func (u UserRepositoryImpl) FindUserById(id uuid.UUID) (dao.User, error) {
	user := dao.User{
		BaseModel: dao.BaseModel{
			ID: id,
		},
	}
	err := u.db.Preload("Permission").First(&user).Error
	if err != nil {
		slog.Error("Got and error when find user by id.", "error", err)
		return dao.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) Save(user *dao.User) (dao.User, error) {
	var err = u.db.Save(user).Error
	if err != nil {
		slog.Error("Got an error when save user.", "error", err)
		return dao.User{}, err
	}
	return *user, nil
}

func (u UserRepositoryImpl) DeleteUserById(id uuid.UUID) error {
	err := u.db.Delete(&dao.User{}, id).Error
	if err != nil {
		slog.Error("Got an error when delete user.", "error", err)
		return err
	}
	return nil
}

func (u UserRepositoryImpl) AddPermissionToUser(userID uuid.UUID, permissionID uuid.UUID) error {
	user := dao.User{
		BaseModel: dao.BaseModel{
			ID: userID,
		},
	}
	permission := dao.Permission{
		BaseModel: dao.BaseModel{
			ID: userID,
		},
	}
	err := u.db.Model(&user).Association("Permissions").Append(&permission)
	if err != nil {
		slog.Error("Got an error when add permission to user.", "error", err)
		return err
	}
	return nil
}

func (u UserRepositoryImpl) DeletePermissionFromUser(userID uuid.UUID, permissionID uuid.UUID) error {
	user := dao.User{
		BaseModel: dao.BaseModel{
			ID: userID,
		},
	}
	permission := dao.Permission{
		BaseModel: dao.BaseModel{
			ID: userID,
		},
	}
	err := u.db.Model(&user).Association("Permissions").Delete(&permission)
	if err != nil {
		slog.Error("Got an error when remove permission from user.", "error", err)
		return err
	}
	return nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&dao.User{})

	return &UserRepositoryImpl{
		db: db,
	}
}
