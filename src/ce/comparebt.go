// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

// Package ce contains go lang coding exercises.
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

//
// https://tour.golang.org/concurrency/8
//
// Exercise: Equivalent Binary Trees
// 1. Implement the Walk function.
//
// 2. Test the Walk function.
//
// The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.
//
// Create a new channel ch and kick off the walker:
//
// go Walk(tree.New(1), ch)
// Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
//
// 3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
//
// 4. Test the Same function.
//
// Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.
//
// The documentation for Tree can be found at: https://godoc.org/golang.org/x/tour/tree#Tree.
//

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	ch2 := make(chan int)
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	var ok1, ok2 bool = true, true
	var v1, v2 int
	for ok1 && ok2 {
		v1, ok1 = <-ch1
		v2, ok2 = <-ch2
		if v1 != v2 {
			return false
		}
	}
	if !ok1 && !ok2 {
		// both are false and no value was found different
		return true
	} else {
		// looks one channel closed earlier than the other, unequal trees
		return false
	}
}

func main() {
	//ch := make(chan int)
	//go Walk(tree.New(1), ch)
	//for j := 0; j < 10; j++ {
	//	fmt.Println(<-ch)
	//}
	s := Same(tree.New(1), tree.New(1))
	t := Same(tree.New(1), tree.New(2))
	st := fmt.Sprintf("s: %t t: %t", s, t)
	fmt.Println(st)
}
