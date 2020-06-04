/*
 * It is a launch point for Binry Tree , BST and AVL Trees. It gives some
 * sample input to validate the functionality of Binary Tree implementation
 * in 'bt' package.
 *
 * MIT License
 * Author: Umesh Patil, Neosemantix, Inc.
 */
package main

import (
	"../../src/bt"
	"fmt"
	"log"
)

// Tests binary tree pyramid printing.
func main() {
	m := make(map[int]int)
	m[0] = 250
	m[1] = 134
	//m[3] = 78
	//bt1 := buildBinaryTree(m)
	//fmt.Println(bt1.Pyramid())

	//Expected ouput:
	//                 250
	//      134                   nil
	// 78         nil        nil        nil

	m[2] = 178
	m[3] = 250
	m[4] = 251
	//bst1 := buildBinarySearchTree(m)
	//fmt.Println(bst1.Pyramid(true))	// include depths for debugging

	// Expected output:
	// 									250
	// 					250                                         251
	//		134                   nil                   nil                   nil
	//	nil        178        nil        nil        nil        nil        nil        nil

	//avlt := buildAvlBstFromMap(m)
	//fmt.Println(avlt.Pyramid(true))	// include depths for debugging

	values := []int{4, 2, 1, 5, 6, 9, 14, 11, 10, 20}
	//}
	avlt := buildAvlBstFromArray(values)
	fmt.Println(avlt.Pyramid(true))
	fmt.Println(avlt.Inorder())

	// Output should look like:

	//																			9(3,4)
	//						5(2,1)                                                                         11(1,2)
	//				2(1,1)                                  6(0,0)                                 10(0,0)                                 14(0,1)
	//		1(0,0)                4(0,0)                nil                 nil                 nil                 nil                 nil                 20(0,0)
	//nil        nil        nil        nil        nil        nil        nil        nil        nil        nil        nil        nil        nil        nil        nil        nil

	//1  2  4  5  6  9  10  11  14  20

}

func buildBinaryTree(val map[int]int) bt.Bt {
	bt2 := bt.NewBt()
	for index := range val {
		fmt.Println(index)
		// we place the value at the specific node/index in the tree
		e := bt2.Insert(index, val[index])
		if e != nil {
			log.Fatal(e)
		}
	}
	return *bt2
}

func buildBinarySearchTree(val map[int]int) bt.Bst {
	bt2 := bt.NewBst()
	for i := 0; i < len(val); i++ {
		fmt.Println(i)
		// unlike binary tree, we all 'insert' values
		// do not place a value at a index
		// rather expect it to be placed at the correct index
		// so as the tree 'remains binary search'; meaning
		// value in the left subtree equal or smaller than root
		// while values in the right subtree are larger.
		// Note that it is not a balanced tree, it is just a
		// Binary Search Tree. So we take first element as the
		// root and push all other elements to Left of the root
		// or to the Right. Nodes with same values of the root
		// are pushed as the immediate Left Child, everything
		// under the current left sub-tree going beneath that.
		e := bt2.Insert(val[i])
		if e != nil {
			log.Fatal(e)
		}
	}
	return *bt2
}

func buildAvlBstFromMap(val map[int]int) bt.Avl {
	avlt := bt.NewAvl()
	for i := 0; i < len(val); i++ {
		e := avlt.Insert(val[i])
		if e != nil {
			log.Fatal(e)
		}
	}
	return *avlt
}

func buildAvlBstFromArray(val []int) bt.Avl {
	avlt := bt.NewAvl()
	for i := 0; i < len(val); i++ {
		e := avlt.Insert(val[i])
		if e != nil {
			log.Fatal(e)
		}
	}
	return *avlt
}
