package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
		buf := bytes.NewBufferString("on\ntwo\nthree\nfour\n")
	
		for {
			line, err := buf.ReadString('\n')
	
			if err != nil {
				if err == io.EOF {
					fmt.Print(line)
					break
				}
				fmt.Println(err)
				break
			}
			fmt.Print(line)
		}


		// slices
		integers := make([]int, 10)

		fmt.Println(integers)

		for i := range integers {
			integers[i] = i
		}

		fmt.Println(integers)

		// map
		sameSalary := map[string]string{"name":"francisco", "job":"web develop", "salary": "500"}

		for key, value := range sameSalary {
			fmt.Println(key + ": " + value )
		}

		// nesting
		numbers := []int{1,2,3}
		letters := []string{"a", "b", "c"}

		for _, number := range numbers {
			for _, letter := range letters {
				fmt.Println(letter)
			}
			fmt.Println(number)
		}

		ints := [][]int {
			[]int{1,2,3},
			[]int{3,2,1},
			[]int{2,1,3},
		}

		for _, v := range ints {
			fmt.Println(v)
		}

		for _, i := range ints {
			for _, j := range i {
				fmt.Println(j)
			}
		}

}