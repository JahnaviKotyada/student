package services

import (
	"stu/models"
	"stu/repository"
)

type StudentService interface {
	GetAllStudents() ([]models.Student, error)
	GetStudentByID(id uint) (*models.Student, error)
	CreateStudent(student *models.Student) error
	UpdateStudent(student *models.Student) error
	DeleteStudent(id uint) error
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) *studentService {
	return &studentService{
		repo: repo,
	}
}

func (s *studentService) GetAllStudents() ([]models.Student, error) {
	return s.repo.GetStudents()
}

func (s *studentService) GetStudentByID(id uint) (*models.Student, error) {
	return s.repo.GetStudentByID(id)
}

func (s *studentService) CreateStudent(student *models.Student) error {
	return s.repo.CreateStudent(student)
}

func (s *studentService) UpdateStudent(student *models.Student) error {
	return s.repo.UpdateStudent(student)
}

func (s *studentService) DeleteStudent(id uint) error {
	return s.repo.DeleteStudent(id)
}
