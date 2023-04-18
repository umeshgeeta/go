/*
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
 * Return the linked list sorted as well.
 *
 * Copyright (c) 2023. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
	"strconv"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1, 1, 2, 3}
	testList(nums)

	nums = []int{}
	testList(nums)

	nums = []int{6}
	testList(nums)

	nums = []int{6, 8, 8}
	testList(nums)

	nums = []int{6, 6}
	testList(nums)

	nums = []int{6, 6, 6, 6, 9, 11, 14, 14, 16}
	testList(nums)
}

func testList(nums []int) {
	list := buildList(nums)
	fmt.Println(printList(list))
	list2 := deleteDuplicates(list)
	fmt.Println(printList(list2))
}

func printList(head *ListNode) string {
	if head == nil {
		return "nil"
	}
	result := ""
	for head != nil {
		result = result + strconv.Itoa(head.Val) + " "
		head = head.Next
	}
	return result
}

func buildList(nums []int) *ListNode {
	var head *ListNode = nil
	var here *ListNode = head
	numCount := len(nums)
	if numCount > 0 {
		ln := ListNode{Val: nums[0], Next: nil}
		head = &ln
		here = head
		for i := 1; i < numCount; i++ {
			node := ListNode{Val: nums[i], Next: nil}
			here.Next = &node
			here = &node
		}
	}
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	start := head
	here := head
	for here != nil {
		previousVal := here.Val
		previousHere := here
		here = here.Next
		for here != nil && here.Val == previousVal {
			here = here.Next
		}
		previousHere.Next = here
	}
	return start
}
