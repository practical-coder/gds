package staff

import (
	"fmt"
	"strings"
)

type employee struct {
	lastName  string
	firstName string
	role      string
	salary    float64
}

type Employee interface {
	SetLastName(lName string)
	SetFirstName(fName string)
	SetRole(r string)
	GetRole() string
	SetSalary(s float64)
	GetSalary() float64
	fmt.Stringer
}

type partTimeEmployee struct {
	employee
	hourlyWage float64
}

type PartTimeEmployee interface {
	Employee
	SetHourlyWage(hourly float64)
	GetHourlyWage() float64
}

func (e *employee) SetSalary(yearly float64) {
	e.salary = yearly
}
func (e employee) GetSalary() float64 {
	return e.salary
}
func (e *employee) SetFirstName(firstName string) {
	e.firstName = firstName
}
func (e employee) GetFirstName() string {
	return e.firstName
}
func (e *employee) SetLastName(lastName string) {
	e.lastName = lastName
}
func (e employee) GetLastName() string {
	return e.lastName
}
func (e *employee) SetRole(role string) {
	e.role = role
}
func (e employee) GetRole() string {
	return e.role
}
func (e employee) String() string {
	b := strings.Builder{}
	b.WriteString("Name: ")
	b.WriteString(fmt.Sprintf("%s %s\n", e.GetFirstName(), e.GetLastName()))
	b.WriteString(fmt.Sprintf("Role: %s\n", e.GetRole()))
	b.WriteString(fmt.Sprintf("Annual salary: $%0.2f\n", e.GetSalary()))
	return b.String()
}

func (pe *partTimeEmployee) SetHourlyWage(amount float64) {
	pe.hourlyWage = amount
}

func (pe partTimeEmployee) GetHourlyWage() float64 {
	return pe.hourlyWage
}

func (pe partTimeEmployee) String() string {
	b := strings.Builder{}
	b.WriteString("Name: ")
	b.WriteString(fmt.Sprintf("%s %s\n", pe.GetFirstName(), pe.GetLastName()))
	b.WriteString(fmt.Sprintf("Role: %s\n", pe.GetRole()))
	b.WriteString(fmt.Sprintf("HourlyWage: $%0.2f\n", pe.GetHourlyWage()))
	return b.String()
}
