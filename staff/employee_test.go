package staff

import (
	"fmt"
	"testing"
)

func TestEmployee(t *testing.T) {
	person := new(employee) // Returns the address of an employee
	person.SetFirstName("Helen")
	person.SetLastName("Rose")
	person.SetRole("Technical Lead")
	person.SetSalary(125_644.0)
	fmt.Println(person)
	hourlyWorker := new(partTimeEmployee) // Returns address
	hourlyWorker.SetFirstName("Mark")
	hourlyWorker.SetLastName("Smith")
	hourlyWorker.SetRole("Software Developer")
	hourlyWorker.SetHourlyWage(85.00)
	fmt.Println(hourlyWorker)
}
