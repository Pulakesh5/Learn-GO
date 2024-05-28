package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	Walk(t1, ch1)
	Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		val1 := <-ch1
		val2 := <-ch2
		if val1 != val2 {
			return false
		}
		//fmt.Println(val)
	}
	return true
}
func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(2), ch)
	defer close(ch)
	//var val int
	for i := 0; i < 10; i++ {
		val := <-ch
		fmt.Println(val)

	}
	Tree1, Tree2 := tree.New(1), tree.New(1)
	check := Same(Tree1, Tree2)
	if check {
		fmt.Println("Two trees are same")
	} else {
		fmt.Println("Two trees are not same")
	}
	fmt.Println(Tree1.String(), "\n", Tree2.String())

	Tree1, Tree2 = tree.New(1), tree.New(2)
	check = Same(Tree1, Tree2)
	if check {
		fmt.Println("Two trees are same")
	} else {
		fmt.Println("Two trees are not same")
	}
	fmt.Println(Tree1.String(), "\n", Tree2.String())
	//binaryTree := tree.New(1)
	//fmt.Println(binaryTree.String())
}
