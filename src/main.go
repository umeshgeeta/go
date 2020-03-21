// author: Umesh Patil
// March 2020
package main

import (
	"./bt"
	"fmt"
	"log"
)

// Tests binary tree pyramid printing.
func main() {
	m := make(map[int]int)
	m[0] = 250
	m[1] = 134
	m[3] = 78
	bt1 := buildBinaryTree(m)
	fmt.Println(bt1.Pyramid())

	//Expected ouput:
	//                 250
	//      134                   nil
	// 78         nil        nil        nil
}

func buildBinaryTree(val map[int]int) bt.Bt {
	bt2 := bt.NewBt()
	for index := range val {
		fmt.Println(index)
		e := bt2.Insert(index, val[index])
		if e != nil {
			log.Fatal(e)
		}
	}
	return *bt2
}
