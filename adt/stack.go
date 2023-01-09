package adt

type Ordered interface {
	~float64 | ~int | ~string
}

type Stack[T Ordered] struct {
	items []T
}

func (stack *Stack[T]) Null() T {
	var result T
	return result
}

func (stack *Stack[T]) Push(item T) {
	if item != stack.Null() {
		stack.items = append(stack.items, item)
	}
}

func (stack *Stack[T]) Pop() T {
	length := len(stack.items)
	if length > 0 {
		result := stack.items[length-1]
		stack.items = stack.items[:(length - 1)]
		return result
	} else {
		return stack.Null()
	}
}

func (stack Stack[T]) Top() T {
	length := len(stack.items)
	if length > 0 {
		return stack.items[length-1]
	} else {
		return stack.Null()
	}
}

func (stack Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}
