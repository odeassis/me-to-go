package main

import "fmt"

func main() {
	items := make(map[string]int)

	items["key1"] = 1
	items["key2"] = 2

	delete(items, "key2")

	fmt.Println(items)
}