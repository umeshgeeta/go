/*
 * https://leetcode.com/problems/symmetric-tree/description/
 *
 * Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
 *
 * The approach is to make BFS - fatten out given binary tree and at each level of the tree check the symmetry.
 * Note how in flattened BFS, nils are populating missing children. Last level is all 'nil' so in itself that
 * does not break the symmetry if the binary tree is all symmetric before that.
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

func main() {
	//vals := []int{1, 2, -1}
	//printSymmetry(vals)
	//
	//vals = []int{1, 2, 2, 3, 4, 4, 3}
	//printSymmetry(vals)
	//
	//vals = []int{1, 2, 2, 3, 4, 4, 3, 11, 12, 12, 13, 13, 12, 12, 11}
	//printSymmetry(vals)

	vals := []int{0, 14, 14, 25, -1, -1, 25, 36, -1, -1, -1, -1, -1, -1, 36}
	printSymmetry(vals)
}

func printSymmetry(vals []int) {
	root := buildTree(vals)
	symmetric := isSymmetric(root)
	fmt.Println(symmetric)
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

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	return walkLevels(bfs(root))
}

func bfs(root *TreeNode) []*TreeNode {
	bfs_array := make([]*TreeNode, 1)
	bfs_array[0] = root
	i := 0
	howManyInLevel := 1
	levelStartIndex := 0
	atLeastOneNonNilChild := true
	done := false
	for i < len(bfs_array) && !done {
		node := bfs_array[i]
		if node != nil {
			ln := node.Left
			rn := node.Right
			bfs_array = append(bfs_array, ln)
			bfs_array = append(bfs_array, rn)
			if ln != nil || rn != nil {
				atLeastOneNonNilChild = true
			}
		} else {
			bfs_array = append(bfs_array, nil)
			bfs_array = append(bfs_array, nil)
		}
		i++
		if howManyInLevel == (i - levelStartIndex) {
			howManyInLevel = 2 * howManyInLevel
			levelStartIndex = i
			if !atLeastOneNonNilChild {
				done = true
			} else {
				// reset for the next iteration
				atLeastOneNonNilChild = false
			}
		}
	}
	return bfs_array
}

func walkLevels(nodes []*TreeNode) bool {
	nodeCount := len(nodes)
	i := 0
	howManyInLevel := 1
	for i+howManyInLevel <= nodeCount {
		j := i + howManyInLevel
		fmt.Printf("i: %d j: %d ", i, j)
		if !symmetricLevel(nodes[i:j]) {
			return false
		}
		howManyInLevel = howManyInLevel * 2
		i = j
	}
	return true
}

func symmetricLevel(levelNodes []*TreeNode) bool {
	count := len(levelNodes)
	fmt.Printf(" count: %d\n", count)
	if count == 1 {
		return true
	}
	if count%2 == 1 {
		// odd nodes, not acceptable
		return false
	}
	i := 0
	j := count - 1
	for j > -1 && i < count-1 && i != j {
		ln := levelNodes[i]
		rn := levelNodes[j]
		if ln == nil {
			if rn != nil {
				return false
			}
			// else both are nil, good
		} else {
			if rn == nil {
				// ln is nil while rn is not
				return false
			}
			// both are non-nil values must be same
			if ln.Val != rn.Val {
				return false
			}
		}
		i++
		j--
	}
	return true
}
