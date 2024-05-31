package main

// List represents a singly-linked list that holds
// values of any type.
import (
	"fmt"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func newList[T any](value T) List[T] {
	newNode := List[T]{nil, value}
	return newNode
}
func addNode[T any](root *List[T], value T) {
	node := root
	for node.next != nil {
		node = node.next
	}
	newNode := newList[T](value)
	node.next = &newNode
}
func main() {
	myList := newList[int](5)

	addNode(&myList, 10)
	addNode(&myList, 15)
	addNode(&myList, 20)

	for temp := &myList; ; {
		fmt.Printf("%v -> ", (*temp).val)
		if (*temp).next != nil {
			temp = (*temp).next
		} else {
			break
		}
	}
}
