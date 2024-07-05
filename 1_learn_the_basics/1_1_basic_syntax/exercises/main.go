package main

import "fmt"

var pscX int = 24
var pscY string = "Bond James"
var pscZ bool = false

type caramelo int

func main() {
	x := 42
	y := "James Bond"
	z := true

	var dog caramelo

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%v - %v - %v\n", pscX, pscY, pscZ)

	s := fmt.Sprintf("%v, %v, %v", pscX, pscY, pscZ)

	fmt.Println(s)

	fmt.Printf("%v - %T\n", dog, dog)

	dog = 42

	fmt.Printf("%v - %T\n", dog, dog)
}