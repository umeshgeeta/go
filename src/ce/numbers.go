// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import "fmt"

func reverse(x int) int {
	result := 0
	num := x
	sign := 1
	if num < 0 {
		sign = -1
		num = sign * num
	}
	divisor := 10
	for num != 0 {
		rem := num % divisor
		num = num / divisor
		result = (result * 10) + rem
	}
	result = result * sign
	var r32 int32 = int32(result)
	var r int = int(r32)
	fmt.Printf("%v %v %v\n", result, r32, r)
	if r != result {
		result = 0
	}
	return result
}

func main() {
	result := reverse(153423646)
	fmt.Printf("Result: %d\n", result)
}
