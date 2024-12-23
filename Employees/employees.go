package Employees

import "fmt"

// Employee interface with GetDetails method
type Employee interface {
	GetDetails() string
}

// FullTimeEmployee struct
type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

// PartTimeEmployee struct
type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

// Implement GetDetails() for FullTimeEmployee
func (e FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Full Time Employee ID: %d, Name: %s, Salary: %d", e.ID, e.Name, e.Salary)
}

// Implement GetDetails() for PartTimeEmployee
func (e PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Part Time Employee ID: %d, Name: %s, Hourly Rate: %d, Hours Worked: %.2f", e.ID, e.Name, e.HourlyRate, e.HoursWorked)
}

// Company struct to manage employees
type Company struct {
	Employees map[string]Employee
}

// AddEmployee adds an employee to the company
func (c *Company) AddEmployee(emp Employee) {
	id := fmt.Sprintf("%d", emp.GetDetails())
	c.Employees[id] = emp
}

// ListEmployees lists all employees in the company
func (c *Company) ListEmployees() {
	for _, emp := range c.Employees {
		fmt.Println(emp.GetDetails())
	}
}
