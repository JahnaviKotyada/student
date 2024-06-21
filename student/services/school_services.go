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

type schoolService struct {
	repo repository.SchoolRepository
}

func NewSchoolService(repo repository.SchoolRepository) *schoolService {
	return &schoolService{
		repo: repo,
	}
}

func (s *schoolService) GetAllSchools() ([]models.School, error) {
	return s.repo.GetSchools()
}

func (s *schoolService) GetSchoolByID(id uint) (*models.School, error) {
	return s.repo.GetSchoolByID(id)
}

func (s *schoolService) CreateSchool(school *models.School) error {
	return s.repo.CreateSchool(school)
}

func (s *schoolService) UpdateSchool(school *models.School) error {
	return s.repo.UpdateSchool(school)
}

func (s *schoolService) DeleteSchool(id uint) error {
	return s.repo.DeleteSchool(id)
}
