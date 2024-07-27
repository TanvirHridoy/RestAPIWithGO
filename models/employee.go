package models

type Employee struct {
	EmployeeId  int    `json:"employeeId" gorm:"primaryKey;column:EmployeeId;autoIncrement"`
	Name        string `json:"name" gorm:"column:Name;type:nvarchar(250);not null"`
	Address     string `json:"address" gorm:"column:Address;type:nvarchar(100);not null"`
	Designation string `json:"designation" gorm:"column:Designation;type:nvarchar(100);not null"`
}

// TableName overrides the table name used by GORM
func (Employee) TableName() string {
	return "Employees"
}
