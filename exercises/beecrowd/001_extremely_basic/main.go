package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	valueA := bufio.NewScanner(os.Stdin)
	valueB := bufio.NewScanner(os.Stdin)
	
	fmt.Print("Valor de A: ")
	valueA.Scan()

	fmt.Print("Valor de B: ")
	valueB.Scan()

	a, errA := strconv.Atoi(valueA.Text())
	b, errB :=strconv.Atoi(valueB.Text()) 

	if errA != nil || errB != nil {
		fmt.Println("Erro na conversão de tipos")
		return
	}

	fmt.Println("X =" , a + b)
}