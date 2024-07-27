package controllers

import (
	"GoAPI/models"
	"GoAPI/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service *services.EmployeeService
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

// CreateEmployee godoc
// @Summary Create a new employee
// @Description Create a new employee with the input payload
// @Tags employees
// @Accept  json
// @Produce  json
// @Param employee body models.Employee true "Create employee"
// @Success 200 {object} models.Employee
// @Router /employees [post]
func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
	var employee models.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateEmployee(&employee); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// GetAllEmployees godoc
// @Summary Get all employees
// @Description Get a list of all employees
// @Tags employees
// @Produce  json
// @Success 200 {array} models.Employee
// @Router /employees [get]
func (c *EmployeeController) GetAllEmployees(ctx *gin.Context) {
	employees, err := c.service.GetAllEmployees()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employees)
}

// GetEmployee godoc
// @Summary Get an employee
// @Description Get an employee by ID
// @Tags employees
// @Produce  json
// @Param id path int true "Employee ID"
// @Success 200 {object} models.Employee
// @Router /employees/{id} [get]
func (c *EmployeeController) GetEmployee(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	employee, err := c.service.GetEmployeeByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// UpdateEmployee godoc
// @Summary Update an employee
// @Description Update an employee with the input payload
// @Tags employees
// @Accept  json
// @Produce  json
// @Param id path int true "Employee ID"
// @Param employee body models.Employee true "Update employee"
// @Success 200 {object} models.Employee
// @Router /employees/{id} [put]
func (c *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var employee models.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee.EmployeeId = id
	if err := c.service.UpdateEmployee(&employee); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Delete an employee by ID
// @Tags employees
// @Produce  json
// @Param id path int true "Employee ID"
// @Success 200 {object} map[string]string
// @Router /employees/{id} [delete]
func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.DeleteEmployee(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
