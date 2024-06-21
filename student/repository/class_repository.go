package repository

import (
	"stu/models"

	"gorm.io/gorm"
)

type ClassRepository interface {
	GetClasses() ([]models.Class, error)
	GetClassByID(id uint) (*models.Class, error)
	CreateClass(class *models.Class) error
	UpdateClass(class *models.Class) error
	DeleteClass(id uint) error
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{
		db: db,
	}
}

func (r *classRepository) GetClasses() ([]models.Class, error) {
	var classes []models.Class
	result := r.db.Find(&classes)
	if result.Error != nil {
		return nil, result.Error
	}
	return classes, nil
}

func (r *classRepository) GetClassByID(id uint) (*models.Class, error) {
	var class models.Class
	result := r.db.First(&class, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &class, nil
}

func (r *classRepository) CreateClass(class *models.Class) error {
	result := r.db.Create(class)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *classRepository) UpdateClass(class *models.Class) error {
	result := r.db.Save(class)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *classRepository) DeleteClass(id uint) error {
	result := r.db.Delete(&models.Class{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
