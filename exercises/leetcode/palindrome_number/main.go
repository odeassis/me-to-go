// Problem url: https://leetcode.com/problems/palindrome-number/

package main

import (
	"fmt"
	"strconv"
)

func isPalidrome(x int) bool {

	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	stringNum := strconv.Itoa(x)

	for i := 0; i <= len(stringNum); i++ {
		if stringNum[i] != stringNum[len(stringNum)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalidrome(10))
}

// 10 -> 01 - false
// 11 -> 11 - true
// 120 -> 021 - false
// 111 -> 111 - true
// 203 -> 302 - false
// 121 -> 121 - true
