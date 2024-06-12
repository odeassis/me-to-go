package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func readContent(file string) []byte {
	content, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return content
}

func main() {
	// bool
	var a bool
	var b bool = false
	c := true

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("a && b:", a&&b)
	fmt.Println("a || b:", a||b)
	fmt.Println("!c:", !c)
	
	var isRainig bool = false
	
	if isRainig {
		fmt.Println("Leve o guarda-chuva")
	} else {
		fmt.Println("Nao precisa levar gurada chuva")
	}

	// int
	var numa int = -20
	var numb = 20
	fmt.Printf("O tipo de %d e %T\n", numa, numa)
	fmt.Printf("O tipo de %d e %T\n", numb, numb)

	var unuma uint16 = 200;
	var inumb int16 = -200;
	fmt.Printf("O tipo de %d e %T\n", unuma, unuma)
	fmt.Printf("O tipo de %d e %T\n", inumb, inumb)

	// float
	var pi = 3.141
	fmt.Printf("O tipo de: %f e: %T\n", pi,pi)

	var price float32 = 102.93
	fmt.Printf("O tipo de: %g e %T\n", price,price)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^X = %.2f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN

	/*
	* Qualquer comparacao com NaN sempre resulta em false
	* (exceto != que e sempre a negacao de ==)
	*/
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // false false false
	
	// complex
	var x complex128 = complex(1, 2)
	fmt.Println("complex128 x:",x)
	
	var y complex128 = complex(3,4)
	fmt.Println("complex128 y:",y)
	
	fmt.Println("complex x * y:", x*y)
	fmt.Println("real x * y:", real(x*y))
	fmt.Println("img x * y:", imag(x*y))

	// byte
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f} // "Hello" em ASCII
	fmt.Printf("%s\n",data)

	content := readContent("1_basic_syntax/1_3_data_types/exemple.txt")

	fmt.Printf("%q\n", content)
	fmt.Printf("%d\n", content)

	// rune
	var ch rune = '世'
	fmt.Printf("Caractere: %c, Codigo Unicode: %U\n", ch, ch)

	str := "Hello, 世界"
	for index, ch := range str {
		fmt.Printf("Indice: %d, Caractere: %c, Codigo Unicode: %U\n", index, ch, ch)
	}

	runes := []rune{'H', 'e', 'l', 'l', 'o', ' ', '世', '界'}
	for _, r := range runes {
		fmt.Printf("%c\n", r)
	}

}