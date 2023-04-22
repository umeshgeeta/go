/*
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

/*
* LeetCode Problem No 101
* https://leetcode.com/problems/symmetric-tree/description/
*
* Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*
* The approach is to make BFS - fatten out given binary tree and at each level of the tree check the symmetry.
* Note how in flattened BFS, nils are populating missing children. Last level is all 'nil' so in itself that
* does not break the symmetry if the binary tree is all symmetric before that.
 */
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

// LeetCode Problem No 104 - Find maximum depth of a binary tree
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

/*
 * LeetCode problem no. 94
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * Given the root of a binary tree, return the inorder traversal of its nodes' values.
 */
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

/*
* LeetCode problem no. 110
* https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
*
* Given an integer array nums where the elements are sorted in ascending order,
* convert it to a height-balanced binary search tree.
 */
func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	switch l {
	case 0:
		return nil
	case 1:
		root := TreeNode{Val: nums[0]}
		return &root
	case 2:
		root := TreeNode{Val: nums[0]}
		rightChild := TreeNode{Val: nums[1]}
		root.Right = &rightChild
		return &root
	case 3:
		root := TreeNode{Val: nums[1]}
		rightChild := TreeNode{Val: nums[2]}
		leftChild := TreeNode{Val: nums[0]}
		root.Right = &rightChild
		root.Left = &leftChild
		return &root
	default:
		middle := l / 2
		root := TreeNode{Val: nums[middle]}
		leftChild := sortedArrayToBST(nums[0:middle])
		rightChild := sortedArrayToBST(nums[middle+1:])
		root.Left = leftChild
		root.Right = rightChild
		return &root
	}
}

/*
 * LeetCode problem no. 111: Balanced Binary Tree
 * https://leetcode.com/problems/balanced-binary-tree/description/
 *
 * Given a binary tree, determine if it is height-balanced
 *
 */
func isBalanced(root *TreeNode) bool {
	b, _ := isBalancedWithMaxDepth(root)
	return b
}

func isBalancedWithMaxDepth(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	if root.Left == nil {
		if root.Right == nil {
			// both children are nil, it is a leaf node
			return true, 1
		}
		// right subtree is not nil
		rb, rd := isBalancedWithMaxDepth(root.Right)
		// left depth is 0, so right depth allowed is max 1
		if rb {
			if rd < 2 {
				// right subtree is balanced & height is no more than 1
				return true, rd + 1
			}
			// else right subtree is balanced, but it's depth is 2 or more
			// since left is empty, it is not allowed
			return false, rd + 1
		}
		// right subtree is not balanced
		return false, rd + 1
	}
	// left is not nil
	lb, ld := isBalancedWithMaxDepth(root.Left)
	if root.Right == nil {
		if lb {
			if ld < 2 {
				// left subtree is balanced & height is no more than 1
				return true, ld + 1
			}
			// else left subtree is balanced, but it's depth is 2 or more
			// since right is empty, not acceptable
			return false, ld + 1
		}
		// left subtree is not balanced
		return false, ld + 1
	}
	// both left & right are not nil
	rb, rd := isBalancedWithMaxDepth(root.Right)
	dd := ld - rd
	if lb && rb && dd < 2 && dd > -2 {
		if ld > rd {
			return true, ld + 1
		}
		return true, rd + 1
	}
	// either diff is larger or left or right subtree is not balanced
	// either ways it is not balanced
	if ld > rd {
		return false, ld + 1
	}
	return false, rd + 1
}

/*
 * LeetCode problem no. 111 - Minimum Depth of Binary Tree
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 *
 * Given a binary tree, find its minimum depth.
 *
 * The minimum depth is the number of nodes along the shortest path from
 * the root node down to the nearest leaf node.
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil {
		mdl := minDepth(root.Left)
		if root.Right != nil {
			// both children are non-nil
			mdr := minDepth(root.Right)
			if mdl < mdr {
				return mdl + 1
			}
			return mdr + 1
		}
		// right is nil, we return left child value which we know is not nil
		return mdl + 1
	}
	// no left child
	if root.Right != nil {
		return minDepth(root.Right) + 1
	}
	// no right child too, it is a leaf node
	return 1
}

/*
 * LeetCode problem no. 112: Path Sum
 * https://leetcode.com/problems/path-sum/description/
 *
 * Given the root of a binary tree and an integer targetSum, return true
 * if the tree has a root-to-leaf path such that adding up all the values
 * along the path equals targetSum.
 *
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root != nil && root.Left == nil && root.Right == nil && root.Val == targetSum {
		// it is a leaf node and has value exact matching targetSum
		return true
	}
	// root is not a leaf node
	// what we need to see if any of the subtrees produce a sum minus root value
	// let us start with Left first
	if hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}
	if hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	// none of the children can produce the exact targetSum
	return false
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

	//vals := []int{1, 2, -1}
	//printSymmetry(vals)
	//
	//vals = []int{1, 2, 2, 3, 4, 4, 3}
	//printSymmetry(vals)
	//
	//vals = []int{1, 2, 2, 3, 4, 4, 3, 11, 12, 12, 13, 13, 12, 12, 11}
	//printSymmetry(vals)

	vals = []int{0, 14, 14, 25, -1, -1, 25, 36, -1, -1, -1, -1, -1, -1, 36}
	printSymmetry(vals)

	root = &TreeNode{4,
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

func printSymmetry(vals []int) {
	root := buildTree(vals)
	symmetric := isSymmetric(root)
	fmt.Println(symmetric)
}
