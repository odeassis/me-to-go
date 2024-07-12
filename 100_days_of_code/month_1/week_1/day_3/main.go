package main

import (
	"fmt"
	"strings"
)

func skewers(barbecue_grill []string) []int {

	vegetarian, non_vegetarian := 0, 0

	for _, v := range barbecue_grill {

		if strings.Contains(v, "x") {
			non_vegetarian++
		} else {
			if strings.Contains(v, "o") {
				vegetarian++
			}
		}
	}

	return []int{vegetarian, non_vegetarian}
}

func main() {
	barbecue_grill_1 := []string{
		"--oooo-ooo--",
		"--xx--x--xx--",
		"--o---o--oo--",
		"--xx--x--ox--",
		"--xx--x--ox--",
	}

	barbecue_grill_2 := []string{
		"--oooo-ooo--",
		"--xxxxxxxx--",
		"--o---",
		"-o-----o---x--",
		"--o---o-----",
	}

	fmt.Println(skewers(barbecue_grill_1))
	fmt.Println(skewers(barbecue_grill_2))
}
