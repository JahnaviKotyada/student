package repository

import (
	"stu/models"

	"gorm.io/gorm"
)

type StudentRepository interface {
	GetStudents() ([]models.Student, error)
	GetStudentByID(id uint) (*models.Student, error)
	CreateStudent(student *models.Student) error
	UpdateStudent(student *models.Student) error
	DeleteStudent(id uint) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{
		db: db,
	}
}

func (r *studentRepository) GetStudents() ([]models.Student, error) {
	var students []models.Student
	result := r.db.Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}

func (r *studentRepository) GetStudentByID(id uint) (*models.Student, error) {
	var student models.Student
	result := r.db.First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

func (r *studentRepository) CreateStudent(student *models.Student) error {
	result := r.db.Create(student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *studentRepository) UpdateStudent(student *models.Student) error {
	result := r.db.Save(student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *studentRepository) DeleteStudent(id uint) error {
	result := r.db.Delete(&models.Student{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
