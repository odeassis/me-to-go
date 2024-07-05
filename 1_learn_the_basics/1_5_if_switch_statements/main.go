package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sum(x int) int {
	return x + rand.Intn(20)
}

func rest(x, y int) (int, int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado do pânico:", r)
		}
	}()

	div := x / y
	rest := x % y

	return div, rest
}

func main() {

	y := 0

	if x := 10; x < y {
		fmt.Println("x < y")
	} else if x == y {
		fmt.Println("x == y")
	} else {
		fmt.Println("x > y")
	}

	if z := sum(y); z > y {
		fmt.Println("z: ", z)
	} else {
		fmt.Println("y:", y)
	}

	if d, r := rest(10, y); y != 0 {
		fmt.Printf("Div: %v - Res: %v\n", d, r)
	} else {
		fmt.Println("Not div for 0")
	}

	scores := map[string]int{"alice": 90, "bob": 76}

	if score, exists := scores["alice"]; exists {
		fmt.Printf("Alice's score is %d\n", score)
	} else {
		fmt.Println("Alice's score is not available")
	}

	fmt.Println(scores["Alice"])

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MAC OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	// Error: the case values are not the same type
	/*
	  number := "1"

	  switch number {
	    case 1:
	      fmt.Println("One")
	    case 2:
	      fmt.Println("Two")
	    default:
	      fmt.Println("Not a number")
	  }
	*/

	today := time.Now().Weekday()

	switch time.Thursday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far way")
	}

	day := time.Now().Weekday()

	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning Sunshine")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}
