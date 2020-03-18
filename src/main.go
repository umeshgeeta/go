// author: Umesh Patil
// March 2020
package main

import (
	"./bt"
	"fmt"
)

// Tests binary tree
func main() {

	bt2 := bt.New(3)
	bt2.Insert(0, 25)
	bt2.Insert(3, 756)
	bt2.Insert(5, 17)
	bt2.Insert(6, 52)
	fmt.Println(bt2.Pyramid())

	bt3 := bt.New(4)
	bt3.Insert(0, 5)
	bt3.Insert(3, 76)
	bt3.Insert(5, 107)
	bt3.Insert(7, 523)
	bt3.Insert(11, 24)
	fmt.Println(bt3.Pyramid())

}
