package main

import (
	"fmt"
	"gds/order"
)

type Student struct {
	ID   int
	Name string
}

func main() {
	students := []any{}
	results := addStudent(students, "Jennifer")
	results = addStudent(results, "Michael")
	results = addStudent(results, "Elaine")
	results = addStudent(results, 159)
	student := Student{ID: 176, Name: "Caroline"}
	structResults := []Student{}
	structResults = addStudent(structResults, student)
	fmt.Println(results)
	fmt.Println(structResults)

	floats := []float64{14.1, 15.7, 16.3, 15.8, 83.2}
	filterFunc := func(item float64) bool {
		return item < 16
	}
	filtered := order.Filter(floats, filterFunc)
	fmt.Println(filtered)
}

func addStudent[T any](students []T, newStudent T) []T {
	return append(students, newStudent)
}
