package main

import "fmt"

func hello(name string) string {
	if name == "" {
		return "Hello, word"
	}

	return "Hello, " + name
}

func main() {
	fmt.Println(hello("Word"))
	var name, age, location = "odeassis", 24, "Brasil"

	println(name, age, location)
}
