package services

import (
	"stu/models"
	"stu/repository"
)

type SchoolService interface {
	GetAllSchools() ([]models.School, error)
	GetSchoolByID(id uint) (*models.School, error)
	CreateSchool(school *models.School) error
	UpdateSchool(school *models.School) error
	DeleteSchool(id uint) error
}

type ClassService interface {
	GetAllClasses() ([]models.Class, error)
	GetClassByID(id uint) (*models.Class, error)
	CreateClass(class *models.Class) error
	UpdateClass(class *models.Class) error
	DeleteClass(id uint) error
}

type StudentService interface {
	GetAllStudents() ([]models.Student, error)
	GetStudentByID(id uint) (*models.Student, error)
	CreateStudent(student *models.Student) error
	UpdateStudent(student *models.Student) error
	DeleteStudent(id uint) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}

// School methods

func (s *service) GetAllSchools() ([]models.School, error) {
	return s.repo.GetSchools()
}

func (s *service) GetSchoolByID(id uint) (*models.School, error) {
	return s.repo.GetSchoolByID(id)
}

func (s *service) CreateSchool(school *models.School) error {
	return s.repo.CreateSchool(school)
}

func (s *service) UpdateSchool(school *models.School) error {
	return s.repo.UpdateSchool(school)
}

func (s *service) DeleteSchool(id uint) error {
	return s.repo.DeleteSchool(id)
}

// Class methods

func (s *service) GetAllClasses() ([]models.Class, error) {
	return s.repo.GetClasses()
}

func (s *service) GetClassByID(id uint) (*models.Class, error) {
	return s.repo.GetClassByID(id)
}

func (s *service) CreateClass(class *models.Class) error {
	return s.repo.CreateClass(class)
}

func (s *service) UpdateClass(class *models.Class) error {
	return s.repo.UpdateClass(class)
}

func (s *service) DeleteClass(id uint) error {
	return s.repo.DeleteClass(id)
}

// Student methods

func (s *service) GetAllStudents() ([]models.Student, error) {
	return s.repo.GetStudents()
}

func (s *service) GetStudentByID(id uint) (*models.Student, error) {
	return s.repo.GetStudentByID(id)
}

func (s *service) CreateStudent(student *models.Student) error {
	return s.repo.CreateStudent(student)
}

func (s *service) UpdateStudent(student *models.Student) error {
	return s.repo.UpdateStudent(student)
}

func (s *service) DeleteStudent(id uint) error {
	return s.repo.DeleteStudent(id)
}
