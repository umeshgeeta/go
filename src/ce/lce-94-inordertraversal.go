/*
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * Given the root of a binary tree, return the inorder traversal of its nodes' values.
 *
 * MIT License
 * Copyright (c) 2023. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	// Visit the left subtree
	result := inorderTraversal(node.Left)

	// Visit the root node
	result = append(result, node.Val)

	// Visit the right subtree
	result = append(result, inorderTraversal(node.Right)...)

	return result
}

func main() {
	root := &TreeNode{4,
		&TreeNode{2,
			&TreeNode{1, nil, nil},
			&TreeNode{3, nil, nil}},
		&TreeNode{6,
			&TreeNode{5, nil, nil},
			&TreeNode{7, nil, nil}},
	}

	fmt.Println("Inorder Traversal:")
	fmt.Println(inorderTraversal(root))
}
