package services

import (
	"stu/models"
	"stu/repository"
)

type ClassService interface {
	GetAllClasses() ([]models.Class, error)
	GetClassByID(id uint) (*models.Class, error)
	CreateClass(class *models.Class) error
	UpdateClass(class *models.Class) error
	DeleteClass(id uint) error
}

type classService struct {
	repo repository.ClassRepository
}

func NewClassService(repo repository.ClassRepository) *classService {
	return &classService{
		repo: repo,
	}
}

func (s *classService) GetAllClasses() ([]models.Class, error) {
	return s.repo.GetClasses()
}

func (s *classService) GetClassByID(id uint) (*models.Class, error) {
	return s.repo.GetClassByID(id)
}

func (s *classService) CreateClass(class *models.Class) error {
	return s.repo.CreateClass(class)
}

func (s *classService) UpdateClass(class *models.Class) error {
	return s.repo.UpdateClass(class)
}

func (s *classService) DeleteClass(id uint) error {
	return s.repo.DeleteClass(id)
}
