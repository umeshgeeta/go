/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * LeetCode problem no. 141: Linked List Cycle
 * https://leetcode.com/problems/linked-list-cycle/description/
 *
 * Given head, the head of a linked list, determine if the linked list has a cycle in it.
 * There is a cycle in a linked list if there is some node in the list that can be reached
 * again by continuously following the next pointer. Internally, pos is used to denote the
 * index of the node that tail's next pointer is connected to. Note that pos is not passed
 * as a parameter.
 *
 * Return true if there is a cycle in the linked list. Otherwise, return false.
 *
 * In the implementation, we use two runners - one fast and one slower. The idea is if
 * there is a cycle, the fast runner will catch the slow runner.
 */
func hasCycle(head *ListNode) bool {
	cycleFound := false
	slowRunner := head
	fastRunner := head
	for !cycleFound && slowRunner != nil && fastRunner != nil {
		if fastRunner.Next == slowRunner {
			cycleFound = true
		}
		slowRunner = slowRunner.Next
		nn := fastRunner.Next
		if nn != nil {
			fastRunner = nn.Next
		}
	}
	return cycleFound
}
