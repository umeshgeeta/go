/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

func main() {

}

func findur(pairs [][]int, n int) int {
	donar := make(map[int]bool)

	pairCount := len(pairs)
	i := 0
	for i < pairCount {
		pair := pairs[i]
		donar[pair[0]] = true
		i++
	}
	if len(donar) == n {
		// everybody is donating, no universal receipt
		return -1
	}
	absent := findAbsentDonar(donar, n)
	// now we need to check if this absent donar receives from everyone
	i = 0
	deletedDonarCount := 0
	for i < pairCount {
		pair := pairs[i]
		if pair[1] == absent {
			// absentee received blood from this donar
			//delete(donar, pair[0])
			deletedDonarCount++
		}
		i++
	}
	if deletedDonarCount == n-1 {
		return absent
	}
	return -1
}

func findAbsentDonar(donar map[int]bool, n int) int {
	i := 1
	for i <= n {
		_, ok := donar[i]
		if !ok {
			return i
		}
		i++
	}
	return -1
}
