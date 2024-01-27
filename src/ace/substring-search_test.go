package ace_test

import (
	"fmt"
	"mymodule/ace"
	"testing"
)

func Test_substringSearch(t *testing.T) {
	result, endingPrefix := ace.SubstringSearch("./../ce/numbers.go", "num")
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Ending Prefix: %v\n", endingPrefix)
}

// func TestSubstringSearch(t *testing.T) {
// 	fileName := "testfile.txt"
// 	pattern := "test"

// 	result, endingPrefix := ace.SubstringSearch(fileName, pattern)

// 	expectedResult := []int{2, 1, 0, 1}
// 	if !intSlicesEqual(result, expectedResult) {
// 		t.Errorf("Expected result %v, but got %v", expectedResult, result)
// 	}

// 	expectedEndingPrefix := []string{"st", "st", "", "st"}
// 	if !stringSlicesEqual(endingPrefix, expectedEndingPrefix) {
// 		t.Errorf("Expected ending prefix %v, but got %v", expectedEndingPrefix, endingPrefix)
// 	}
// }

// func intSlicesEqual(slice1, slice2 []int) bool {
// 	if len(slice1) != len(slice2) {
// 		return false
// 	}
// 	for i := range slice1 {
// 		if slice1[i] != slice2[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func stringSlicesEqual(slice1, slice2 []string) bool {
// 	if len(slice1) != len(slice2) {
// 		return false
// 	}
// 	for i := range slice1 {
// 		if slice1[i] != slice2[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
