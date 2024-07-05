// Poblem url: https://leetcode.com/problems/two-sum/

package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	result := make([]int, 2)

	for i, v := range nums[:len(nums)-1] {
		for j, jv := range nums[i+1:] {
			if v+jv == target {
				result[0] = i
				result[1] = j + 1 + i
				break
			}
		}
	}

	return result
}

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(TwoSum(nums, target))
}
