package repositories

import (
	"GoAPI/models"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Create(employee *models.Employee) error {
	return r.DB.Create(employee).Error
}

func (r *EmployeeRepository) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Find(&employees).Error
	return employees, err
}

func (r *EmployeeRepository) GetByID(id int) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.First(&employee, id).Error
	return &employee, err
}

func (r *EmployeeRepository) Update(employee *models.Employee) error {
	return r.DB.Save(employee).Error
}

func (r *EmployeeRepository) Delete(id int) error {
	return r.DB.Delete(&models.Employee{}, id).Error
}
