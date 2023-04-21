/*
 * MIT License
 * Copyright (c) 2023. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// root is not nil here
	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)
	// we add 1 of the root node
	if ld > rd {
		return ld + 1
	}
	return rd + 1
}

//
// Common code to be refactored
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(values []int) *TreeNode {
	return connectNodes(createNodes(values))
}

func createNodes(values []int) []*TreeNode {
	count := len(values)
	tns := make([]*TreeNode, count)
	i := 0
	for i < count {
		v := values[i]
		if v >= 0 {
			tn := TreeNode{Val: v}
			tns[i] = &tn
		} else {
			// it is -ve
			tns[i] = nil
		}
		i++
	}
	return tns
}

func connectNodes(tns []*TreeNode) *TreeNode {
	i := 1
	nodeCount := len(tns)
	for i < nodeCount {
		if tns[i] != nil {
			pi := i / 2
			r := i % 2
			if r == 0 {
				// even, right node
				parent := tns[pi-1]
				if parent != nil {
					parent.Right = tns[i]
				}
			} else {
				// add
				parent := tns[pi]
				if parent != nil {
					parent.Left = tns[i]
				}
			}
		}
		// else it is nil, nothing to attach to the parent
		i++
	}
	return tns[0]
}

func main() {
	vals := []int{1, 2, -1}
	root := buildTree(vals)
	fmt.Printf("max depth: %d\n", maxDepth(root))

	vals = []int{1, 2, 2, 3, 4, 4, 3}
	root = buildTree(vals)
	fmt.Printf("max depth: %d\n", maxDepth(root))

	vals = []int{1, 2, 2, 3, 4, 4, 3, 11, 12, 12, 13, 13, 12, 12, 11}
	root = buildTree(vals)
	fmt.Printf("max depth: %d\n", maxDepth(root))

	vals = []int{0, 14, 14, 25, -1, -1, 25, 36, -1, -1, -1, -1, -1, -1, 36}
	root = buildTree(vals)
	fmt.Printf("max depth: %d\n", maxDepth(root))
}
