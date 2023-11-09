package main

import (
	"crypto/sha256"
	"fmt"
)

func reverse(copy [3]int) {
	for i, j := 0, len(copy)-1; i < j; i, j = i+1, j-1{
		copy[i], copy[j] = copy[j], copy[i]
	}
	fmt.Printf("Reversed array: %v\n", copy)
}

func one(prt *[32]byte) {
	for i := range prt {
		prt[i] = 1
	}
}

func zero(prt *[32]byte) {
	*prt = [32]byte{}
}

func main() {
	// var array_variable = [size]datatype{elemente of array}
	var ages = [5]int{17,18,15,19,14}

	fmt.Println(ages) // -> [17 18 15 19 14]

	// Accessing arrays elements
	fullName := []string{"francisco", "de", "assis"}

	fmt.Println(fullName[0]) // -> francisco
	fmt.Println(fullName[2]) // -> assis

	// Initialize an array
	var integers[3] int
	integers[0] = 9
	integers[1] = 2
	integers[2] = 5

	integersTwo := [5]int{0:8, 3:4}
	fmt.Println(integersTwo) // -> [8 0 0 4 0]

	// Changing the array elements
	weather := [3]string{"Rainy", "Sunny", "Cloudy"}

	weather[1] = "Stromy"
	fmt.Println(weather) // -> [Rainy Stromy Cloudy]

	// Find the lenght of an array
	var arrayOfIntegers = [...]int{1,2,3,4,5,6,7}
	length := len(arrayOfIntegers)

	fmt.Println(length) // -> 7

	// Looping
	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}

	// Multidimensional array
	arrayInteger := [2][2]int{{1,2},{3,4}}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arrayInteger[i][j])
		}
	}

	// Passing arrays to functions
	oArray := [3]int{1,2,3}
	reverse(oArray)
	fmt.Printf("Original array: %v\n", oArray)


	c1 := sha256.Sum256([]byte("a"))
  c2 := sha256.Sum256([]byte("A"))

  fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	zeroTest := [32]byte{1,2,3,4,5,6,7,8,9,10}
	one(&zeroTest)
	fmt.Println(zeroTest)
	zero(&zeroTest)
	fmt.Println(zeroTest)

	
  // Exercise
  exercise1 := make([]int, 3, 10)
  fmt.Println(arrayAverage(exercise1))
}