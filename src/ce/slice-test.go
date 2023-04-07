/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 2)
	s = append(s, 2, 6, 8)
	fmt.Println(s)

	var ss []int
	ss = make([]int, 2)
	ss = append(ss, 2)

	fmt.Println(ss)

	// Output:
	//
	// [0 0 2 6 8]
	// [0 0 2]
	//
	// So append always "adds" new capacity & elements at the end
	// if you want to use existing capacity, you can refer to index.
}
