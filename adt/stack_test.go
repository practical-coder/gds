package adt

import (
	"fmt"
	"testing"
)

func TestStringStack(t *testing.T) {
	nameStack := Stack[string]{}
	nameStack.Push("Zuza")
	nameStack.Push("Ada")
	topOfStack := nameStack.Top()
	if topOfStack != nameStack.Null() {
		fmt.Printf("\nTop of the stack is %v\n", topOfStack)
	}
	poppedFromStack := nameStack.Pop()
	if poppedFromStack != nameStack.Null() {
		fmt.Printf("\nValue popped from stack is %v\n", poppedFromStack)
	}
	poppedFromStack = nameStack.Pop()
	if poppedFromStack != nameStack.Null() {
		fmt.Printf("\nValue popped from stack is %v\n", poppedFromStack)
	}
}
