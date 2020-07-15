// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

// Slice exercise from Go Tour: https://tour.golang.org/moretypes/18
// The pic.Show prints pretty picture on the tour web page. here is simply
// prints Hex strRep.
package main

import (
	"golang.org/x/tour/pic"
	"strings"
)

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i, _ := range result {
		result[i] = make([]uint8, dx)
		for j, _ := range result[i] {
			result[i][j] = uint8(i * j)
		}
	}
	return result
}

// count the words in the input strRep
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
	for _, aWord := range words {
		if count, ok := result[aWord]; ok {
			result[aWord] = count + 1
		} else {
			result[aWord] = 1
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}
