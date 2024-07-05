package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("value: %v - type: %T\n", i, i)
	fmt.Printf("value: %v - type: %T\n", f, f)
	fmt.Printf("value: %v - type: %T\n", u, u)

	var s interface{} = "Hello"

	// converte interface{} para string
	ch := s.(string)

	fmt.Printf("value: %s - type: %T\n", s, s)
	fmt.Printf("value: %s - type: %T\n", ch, ch)

	// converte interface{} para string com verificacao de seguranca
	if ch, ok := s.(string); ok {
		fmt.Println(ch)
	} else {
		fmt.Println("nao e uma string")
	}

	a := strconv.Itoa(300)
	fmt.Printf("%q\n", a)

	user := "odeassis"
	lines := 30
	points := 102.30

	fmt.Println("Congratulations, " + user + "! You just wrote " + strconv.Itoa(lines) + " lines of code.")

	fmt.Println(fmt.Sprint("Congratulations, ", user, "! You winner ", points, " XP points."))

	xp_yesterday := "230.39"
	xp_today := "102"

	yesterday, err := strconv.ParseFloat(xp_yesterday, 64)
	if err != nil {
		log.Fatal(err)
	}

	totay, err := strconv.Atoi(xp_today)
	if err != nil {
		log.Fatal(err)
	}

	boots_xp := yesterday * float64(totay/lines)

	fmt.Println("xp boots: ", boots_xp)

	x := "my string"

	b := []byte(x)

	c := string(b)

	fmt.Println(x)

	fmt.Println(b)

	fmt.Println(c)
}
