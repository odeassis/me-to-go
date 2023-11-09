# O básico da sintaxe

## Comentários

- Comentários de linha única começam com `//`
- Comentários de várias linhas são delimitados por `/*` no inicio e `*/` no final

  ```go
  // Este é um comentário de linha única

  /*
    Este é um comentário
    de várias linhas
  */
  ```

## Pacotes

- Um programa em Go é organizado em pacotes
- O pacote principal é chamado de `main` e é o ponto de entrada de um programa executável

  ```go
    package main
  ```

## Importações

- Para usar funcionalidades de outros pacotes, você precisa importá-los

  ```go
    import "fmt"
  ```

## Funções

- Funções são blocos de código que podem ser chamados
- A função `main` é a primeira função executada em um programa Go

  ```go
    func main() {
      // Seu codigo aqui
    }
  ```

## Variáveis

- Variáveis são declaradas usando a sintaxe `var nome tipo`
- Go é uma linguagem com tipagem estática, o que significa que você deve especificar o tipo da variável

  ```go
    var x int
    x = 10

    var nome string = "odeassis"
  ```

### Declaração curta de variáveis

- Você pode usar `:=` para declarar e atribuir valores a variáveis de forma mais consisa

  ```go
    firstName := "Francisco"
    age := "24"
  ```

## Estrutura de controle

- Go oferece estruturas de controle padrão, como `if`, `for` e `switch`

  ```go
    if age >= 18 {
      fmt.Println("Maior de idade")
    }


    for i := 0; i < 5; i++ {
      fmt.Println(i)
    }

    weekDay := "quarta"

    switch weekDay {
      case "segunda":
        fmt.Println("Dia de trabalho")
      case "sabado", "domingo":
        fmt.Println("Fim de semana")
      default:
        fmt.Println("Outro dia")
    }
  ```

## Arrays e Slices

- Arrays têm um tamanho fixo, enquanto slices são flexíveis
- Para criar um slice, você pode usar a função `make`

  ```go
    var array [5]int
    slice := []int{1,2,3,4,5}

    // Criar um slice de inteiros com tamanho inicial de 3 e capacidade de 5
    mySlice := make([]int,3,5)

    mySlice[0] = 1
    mySlice[1] = 2
    mySlice[2] = 0

    fmt.Println("Slice: ", mySlice)
    fmt.Println("Tamanho:", len(mySlice))
    fmt.Println("Capacidade:", cap(mySlice))
  ```

## Estrutura de dados

- Go suporta structs para definir tipos de dados personalizados

  ```go
    type Pessoa struct {
      name string
      age int
    }
  ```

## Funções personalizadas

- Você pode criar suas próprias funções personalizadas

  ```go
    func sum(a, b int) int {
      return a + b
    }
  ```
