package main

import "fmt"

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
	filtered := GenFilter(floats, func(item float64) bool {
		return item < 16
	})
	fmt.Println(filtered)
}

func addStudent[T any](students []T, newStudent T) []T {
	return append(students, newStudent)
}

func GenFilter[T any](input []T, f func(T) bool) []T {
	results := make([]T, 0)
	for _, item := range input {
		if f(item) {
			results = append(results, item)
		}
	}
	return results
}