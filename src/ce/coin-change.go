// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	cc := len(coins)
	coinCount := make([][]int, amount+1)
	for i := range coinCount {
		coinCount[i] = make([]int, cc+1)
	}
	for i := 0; i < cc+1; i++ {
		coinCount[0][i] = 0
	}
	fmt.Printf("%v\n", coinCount[0])
	for amt := 1; amt < amount+1; amt++ {
		minCoinCount := 0
		for coin := 0; coin < cc; coin++ {
			coinValue := coins[coin]
			if amt >= coinValue {
				lessAmt := amt - coinValue
				coinCount[amt][coin] = coinCount[lessAmt][cc] + 1
			} else {
				coinCount[amt][coin] = 0
			}
			if minCoinCount == 0 {
				minCoinCount = coinCount[amt][coin]
			} else {
				if coinCount[amt][coin] > 0 && coinCount[amt][coin] < minCoinCount {
					minCoinCount = coinCount[amt][coin]
				}
				// else minimum does not change
			}
		}
		coinCount[amt][cc] = minCoinCount
		fmt.Printf("%v\n", coinCount[amt])
	}
	return coinCount[amount][cc]
}

func main() {
	//coins := []int{1,2,5}
	//fmt.Printf("%d\n", coinChange(coins, 11))
	coins := []int{2}
	fmt.Printf("%d\n", coinChange(coins, 3))
}
