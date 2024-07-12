package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findNemo(quote string) int {
	words := strings.Fields(quote)

	for i, v := range words {
		if v == "Nemo" {
			return i + 1
		}
	}

	return 0

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Write a quote: ")
	input, _ := reader.ReadString('\n')

	idxNemo := findNemo(input)

	if idxNemo == 0 {
		fmt.Println("I can't find Nemo :(")
	} else {
		fmt.Printf("I found Nemo at %v!\n", idxNemo)
	}
}
