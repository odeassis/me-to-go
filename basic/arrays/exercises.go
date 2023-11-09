package main

// Exercise 1
// Write a function that takes an array of integers and returns the average of the array.
func arrayAverage(array []int) float64 {
	var sum int

	for _, value := range array {
		sum += value
	}

	average := float64(sum)/float64(len(array))

	return average
}


// Exercise 2
// Write a function that takes an array of integers and returns the sum of the array.

// Exercise 3
// Write a function that takes an array of integers and returns the largest number in the array.

// Exercise 4
// Write a function that takes an array of integers and returns the smallest number in the array.

// Exercise 5
// Write a function that takes an array of integers and returns the index of the largest number in the array.

// Exercise 6
// Write a function that takes an array of integers and returns the index of the smallest number in the array.

// Exercise 7
// Write a function that takes an array of integers and returns the number of even numbers in the array.

// Exercise 8
// Write a function that account numbers of bits that are different between two SHA256 hashes.
