package repository

import (
	"stu/models"

	"gorm.io/gorm"
)

type Repository interface {
	// School methods
	GetSchools() ([]models.School, error)
	GetSchoolByID(id uint) (*models.School, error)
	CreateSchool(school *models.School) error
	UpdateSchool(school *models.School) error
	DeleteSchool(id uint) error

	// Class methods
	GetClasses() ([]models.Class, error)
	GetClassByID(id uint) (*models.Class, error)
	CreateClass(class *models.Class) error
	UpdateClass(class *models.Class) error
	DeleteClass(id uint) error

	// Student methods
	GetStudents() ([]models.Student, error)
	GetStudentByID(id uint) (*models.Student, error)
	CreateStudent(student *models.Student) error
	UpdateStudent(student *models.Student) error
	DeleteStudent(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// School methods

func (r *repository) GetSchools() ([]models.School, error) {
	var schools []models.School
	result := r.db.Find(&schools)
	if result.Error != nil {
		return nil, result.Error
	}
	return schools, nil
}

func (r *repository) GetSchoolByID(id uint) (*models.School, error) {
	var school models.School
	result := r.db.First(&school, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &school, nil
}

func (r *repository) CreateSchool(school *models.School) error {
	result := r.db.Create(school)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) UpdateSchool(school *models.School) error {
	result := r.db.Save(school)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) DeleteSchool(id uint) error {
	result := r.db.Delete(&models.School{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Class methods

func (r *repository) GetClasses() ([]models.Class, error) {
	var classes []models.Class
	result := r.db.Find(&classes)
	if result.Error != nil {
		return nil, result.Error
	}
	return classes, nil
}

func (r *repository) GetClassByID(id uint) (*models.Class, error) {
	var class models.Class
	result := r.db.First(&class, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &class, nil
}

func (r *repository) CreateClass(class *models.Class) error {
	result := r.db.Create(class)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) UpdateClass(class *models.Class) error {
	result := r.db.Save(class)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) DeleteClass(id uint) error {
	result := r.db.Delete(&models.Class{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Student methods

func (r *repository) GetStudents() ([]models.Student, error) {
	var students []models.Student
	result := r.db.Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}

func (r *repository) GetStudentByID(id uint) (*models.Student, error) {
	var student models.Student
	result := r.db.First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

func (r *repository) CreateStudent(student *models.Student) error {
	result := r.db.Create(student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) UpdateStudent(student *models.Student) error {
	result := r.db.Save(student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) DeleteStudent(id uint) error {
	result := r.db.Delete(&models.Student{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
