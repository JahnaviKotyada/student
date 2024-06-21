package repository

import (
	"stu/models"

	"gorm.io/gorm"
)

type SchoolRepository interface {
	// School methods
	GetSchools() ([]models.School, error)
	GetSchoolByID(id uint) (*models.School, error)
	CreateSchool(school *models.School) error
	UpdateSchool(school *models.School) error
	DeleteSchool(id uint) error
}

type schoolRepository struct {
	db *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) SchoolRepository {
	return &schoolRepository{
		db: db,
	}
}

// School methods

func (r *schoolRepository) GetSchools() ([]models.School, error) {
	var schools []models.School
	result := r.db.Find(&schools)
	if result.Error != nil {
		return nil, result.Error
	}
	return schools, nil
}

func (r *schoolRepository) GetSchoolByID(id uint) (*models.School, error) {
	var school models.School
	result := r.db.First(&school, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &school, nil
}

func (r *schoolRepository) CreateSchool(school *models.School) error {
	result := r.db.Create(school)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *schoolRepository) UpdateSchool(school *models.School) error {
	result := r.db.Save(school)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *schoolRepository) DeleteSchool(id uint) error {
	result := r.db.Delete(&models.School{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
