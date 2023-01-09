package main

import (
	"fmt"
	"gds/words"
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

	//blackjack.Play()
	words.Permutations("parties")
}

func addStudent[T any](students []T, newStudent T) []T {
	return append(students, newStudent)
}
