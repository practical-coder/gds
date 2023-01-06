package order

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	a := []int{3, 2, 4, 6, 5,8,7,12,9,11,10}
	fmt.Println(a)
	Bubble[int](a)
	fmt.Println(a)
}
