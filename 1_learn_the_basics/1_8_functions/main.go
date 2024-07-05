package main

import "fmt"

func main() {

	// total := func() int {
	// 	return mult_sum(1,2,3,4,5)
	// }()

	// fmt.Println(total)

	numbers := []int{4,5,6}
	generateCombinations([]int{1,2,3}, numbers)
}

func sum(a int, b int) int {
	return a + b
}

func sub(a, b int) (z int) {
	z = a - b;
	return
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}

func max(a, b int) (int, bool) {
	if a > b {
		return a, true
	}
	return b, false
}

func mult_sum(numbers ...int) int {
	var total int

	for _, num := range numbers {
		total += num
	}

	return total
}

func generateCombinations(current []int, remaining []int) {
	if len(remaining) == 0 {
		fmt.Println(current)
		return
	}

	generateCombinations(current, remaining[1:])
	
	generateCombinations(append(current, remaining[0]), remaining[1:])
}
