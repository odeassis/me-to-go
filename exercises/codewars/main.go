package main

import "fmt"

/*
O código contém dois loops aninhados. O primeiro loop percorre todos os elementos de arr
uma vez, equanto o segundo loop itera sobre os elementos restantes a partir da posição atual
do primeiro loop.

Considerando o pior caso, onde nenhum elemento é menor que os outros elementos à sua direita.
Nesse caso, o segundo loop será executado completamente em cada iteração do primeiro loop.
*/
func Smaller(arr []int) []int {

	// Cria um slice vazia com capacidade igual ao tamanho de arr. Isso tem complexidade O(1).
	result := make([]int, 0, len(arr))
	smallerCounter := 0

	// O primeiro loop executa len(arr) interações.
	for i := 0; i < len(arr); i++ {
		// O segundo loop executa len(arr) - (i + 1) iterações em cada iteração do primeiro loop.
		for k := i + 1; k < len(arr); k++ {
			// A condição if é verificada len(arr) - (i + 1) vezes. Caso a condição seja veradeira, a
			// variável smallerCounter é incrementada. Essa operação tem complexidade O(1).
			if arr[i] > arr[k] {
				smallerCounter += 1
			} else {
				continue
			}
		}
		// Realiza uma operação de anexar um elemento à slice result. Em geral, a operação de anexar
		// tem complexidade O(1). No entanto, pode ser O(n) em casos raros quando a capacidade do slice
		// precisa ser aumentada.
		result = append(result, smallerCounter)
		smallerCounter = 0
	}

	return result
}

/*
  Portanto, a complexidade total do código é dada por:
  O(1) + O(n) * O(n) * O(1) + O(n) * (O(1) + O(n)) + O(1)

    - O(1) -> result := make([]int, 0, len(arr)) e smallerCounter = 0.
    - O(n) -> primeiro loop for.
    - O(n) * O(n) * O(1) -> loop aninhados. O primeiro loop executa n interações e, em cada interação
      o segundo loop executa n - (i + 1) interações. Multiplicando essas duas quantidades, temos n * (n - (i + 1))
      que simpifica pra n * n. O termo * O(1) indica que as operações dentro dos loops aninhados têm complexidade constante
    - O(n) * (O(1) + O(n)) -> operações dentro do segundo loop. A condição if é uma operação de tempo constante, enquanto
      o incremento de smallerCounter é uma operação de tempo constante também. Portanto a complexidade total dentro do
      segundo loop é O(1) + O(n), pois estamos considerando o pior caso em que o loop é execultado completamenete.
    - O(1) -> return result é uma operação de tempo constante, pois retorna a slice result.

  Simplificando:

  O(1) + O(n^2) + O(n) + O(n^2) + O(1)

  A complexidade dominante é O(n^2), pois o termo O(n^2) é o de maior ordem. Portanto, a complexidade de tempo do código é
  O(n^2), onde n é o tamanho da entrada arr.
*/

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := []int{5, 4, 3, 2, 1}

	fmt.Println(Smaller(arr))
	fmt.Println(Smaller(arr2))
}
