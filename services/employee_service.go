package services

import (
	"GoAPI/models"
	"GoAPI/repositories"
)

type EmployeeService struct {
	repo *repositories.EmployeeRepository
}

func NewEmployeeService(repo *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) error {
	return s.repo.Create(employee)
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.repo.GetAll()
}

func (s *EmployeeService) GetEmployeeByID(id int) (*models.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *EmployeeService) UpdateEmployee(employee *models.Employee) error {
	return s.repo.Update(employee)
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	return s.repo.Delete(id)
}
