// Exibe os argumentos da linha de comando
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}

	//for _,arg := range os.Args[1:] {
	//  s += sep + arg
	//  sep = ""
	//}

	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], ""))
}
