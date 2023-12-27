package repository

import (
	"auth-backend/app/domain/dao"
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type DepartmentRepository interface {
	FindAllDepartments() ([]dao.Department, error)
	FindDepartmentById(id uuid.UUID) (dao.Department, error)
	Save(Department *dao.Department) (dao.Department, error)
	DeleteDepartmentById(id uuid.UUID) error
}

type DepartmentRepositoryImpl struct {
	db *gorm.DB
}

func (r DepartmentRepositoryImpl) FindAllDepartments() ([]dao.Department, error) {
	var Departments []dao.Department

	var err = r.db.Preload("Departments").Find(&Departments).Error
	if err != nil {
		slog.Error("Got an error finding all couples.", "error", err)
		return nil, err
	}

	return Departments, nil
}

func (r DepartmentRepositoryImpl) FindDepartmentById(id uuid.UUID) (dao.Department, error) {
	Department := dao.Department{
		BaseModel: dao.BaseModel{
			ID: id,
		},
	}
	err := r.db.Preload("Departments").First(&Department).Error
	if err != nil {
		slog.Error("Got and error when find Department by id.", "error", err)
		return dao.Department{}, err
	}
	return Department, nil
}

func (r DepartmentRepositoryImpl) Save(Department *dao.Department) (dao.Department, error) {
	var err = r.db.Save(Department).Error
	if err != nil {
		slog.Error("Got an error when save Department.", "error", err)
		return dao.Department{}, err
	}
	return *Department, nil
}

func (r DepartmentRepositoryImpl) DeleteDepartmentById(id uuid.UUID) error {
	err := r.db.Delete(&dao.Department{}, id).Error
	if err != nil {
		slog.Error("Got an error when delete Department.", "error", err)
		return err
	}
	return nil
}

func DepartmentRepositoryInit(db *gorm.DB) *DepartmentRepositoryImpl {
	db.AutoMigrate(&dao.Department{})
	return &DepartmentRepositoryImpl{
		db: db,
	}
}