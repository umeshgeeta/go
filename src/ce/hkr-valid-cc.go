/*
 * https://www.hackerrank.com/challenges/validating-credit-card-number/problem
 *
 * You and Fredrick are good friends. Yesterday, Fredrick received  credit cards from ABCD Bank.
 * He wants to verify whether his credit card numbers are valid or not. You happen to be great at regex so
 * he is asking for your help!
 *
 * A valid credit card from ABCD Bank has the following characteristics:
 * ► It must start with a 4, 5 or 6.
 * ► It must contain exactly 16 digits.
 * ► It must only consist of digits (0-9).
 * ► It may have digits in groups of 4, separated by one hyphen "-".
 * ► It must NOT use any other separator like '  ' , '_', etc.
 * ► It must NOT have  4 or more consecutive repeated digits.
 */
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func validateCreditCardNumber(cardNumber string) bool {
	// Must start with 4, 5 or 6
	if !regexp.MustCompile(`^[456]`).MatchString(cardNumber) {
		return false
	}

	// Must only consist of digits (0-9) and optional hyphens
	if !regexp.MustCompile(`^[\d\-]+$`).MatchString(cardNumber) {
		return false
	}

	// Must contain exactly 16 digits
	cardNumber2 := strings.ReplaceAll(cardNumber, "-", "")
	if len(cardNumber2) != 16 {
		return false
	}

	if cardNumber2 < cardNumber {
		// It means it has hyphens. If so, those must be in allowed patter
		// We want to allow maximum of 3 hyphens...
		if !regexp.MustCompile(`^\d{4}-\d{4}-\d{4}-\d{4}$`).MatchString(cardNumber) {
			return false
		}
		// Even if hyphens are 3 or less, we still need to check for consecutive digits
		// not more than 4; hence we carry forward the result here.
	}

	// Must NOT have 4 or more consecutive repeated digits
	for i := 0; i < len(cardNumber)-3; i++ {
		if cardNumber[i] == cardNumber[i+1] && cardNumber[i+1] == cardNumber[i+2] && cardNumber[i+2] == cardNumber[i+3] {
			return false
		}
	}

	return true
}

func main() {
	cardNumbers := []string{
		"49927398716",           // false
		"4992-7398-716",         // false
		"1234_5678_9012_3456",   // false
		"4444444444444444",      // false
		"5555555555554444",      // false
		"2223003122003222",      // false
		"4253625879615786",      // true
		"4424424424442444",      // true
		"5122-2368-7954-3214",   // true
		"5122-2368-7954 - 3214", // false
		"0525362587961578",      // false
		"42536258796157867",     // false
		"4424444424442444",      // false
		"44244x4424442444",      // false
	}

	for _, cardNumber := range cardNumbers {
		valid := validateCreditCardNumber(cardNumber)
		fmt.Printf("%s is valid: %t\n", cardNumber, valid)
	}
}
