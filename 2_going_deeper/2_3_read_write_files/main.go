package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Create
	f, createError := os.Create("file.txt")

	if createError != nil {
		panic(createError)
	}

	// Write

	_, writeError := f.WriteString("Salve o Corinthians, O campeão dos campeões...")

	if writeError != nil {
		panic(writeError)
	}

	f.Close()

	// Read
	file, readError := os.ReadFile("file.txt")

	if readError != nil {
		panic(readError)
	}

	fmt.Println(file)
	fmt.Printf("%q\n", file)

	corinthians, err := os.Open("hino.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(corinthians)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

}
