package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const daysOfYear = 365

func calDays(age string) int64 {

	input := strings.TrimRight(age, "\r\n")

	f_age, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		panic(err)
	}

	return f_age * daysOfYear
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')

	days := calDays(age)

	fmt.Println(days)

}
