package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convert(s_miles []string) []int {
	var outcome []int

	for _, v := range s_miles {
		if n, err := strconv.Atoi(v); err != nil {
			continue
		} else {
			outcome = append(outcome, n)
		}
	}

	return outcome
}

func progressDays(miles []int) int {
	var progress int

	fmt.Println(miles)

	for i, v := range miles[:len(miles)-1] {

		if v < miles[i+1] {
			progress++
		}
	}

	return progress
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a_miles []string

	fmt.Println("Enter with your miles. ex: 1,2,10,3")

	if miles, err := reader.ReadString('\n'); err != nil {
		panic(err)
	} else {
		a_miles = strings.Split(strings.TrimRight(miles, "\r\n"), ",")
	}

	fmt.Println(progressDays(convert(a_miles)))
}
