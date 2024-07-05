package main

import "fmt"

/*
	Você é o desenvolvedor de um sistema de gerenciamento de bibliotecas e precisa
	criar uma Struct para representar um livro. Cada livro tem um título, autor,
	ano de publicação, ISBN (um número único de 13 dígitos) e uma lista de exemplares.

	**Requisitos:**

	1. Crie uma Struct chamada `Book` com os campos mencionados acima.
	2. O campo `ISBN` deve ser uma string de 13 caracteres.
	3. O campo `exemplares` deve ser um slice de structs `Exemplar`.
	4. A Struct `Exemplar` deve ter os campos `id` (um inteiro único),
		`localizacao`(uma string representando a localização do exemplar na biblioteca) e
		`disponivel` (um booleano indicando se o exemplar está disponível para empréstimo).

	**Desafios:**

	1. Implemente a Struct `Book` e a Struct `Exemplar` de acordo com os requisitos.
	2. Adicione um método `AdicionarExemplar` à Struct `Book` que permite adicionar um novo exemplar ao livro.
	3. Adicione um método `RemoverExemplar` à Struct `Book` que permite remover um exemplar do livro baseado em seu ID.
*/

type Book struct {
    Titulo string
    Autor string
    AnoPub int
    ISBN string
    Exemplares []Exemplar
}

type Exemplar struct {
    ID int
    Localizacao string
    Disponivel bool
}

func (b *Book) adicionarExemplar(exemplar Exemplar) {
    b.Exemplares = append(b.Exemplares, exemplar)
}

func (b *Book) removerExemplar(id int) {
    for i,v := range b.Exemplares {
        if v.ID == id {
            b.Exemplares = append(b.Exemplares[:i], b.Exemplares[i+1:]...)
        }
    }
}

func main() {
    // Criar um livro
    livro := Book{
        Titulo:  "O Senhor dos Anéis",
        Autor:   "J.R.R. Tolkien",
        AnoPub:  1954,
        ISBN:   "9788535915378",
        Exemplares: []Exemplar{
            {
                ID: 1,
                Localizacao: "Estante 1",
                Disponivel: true,
            },
            {
                ID: 2,
                Localizacao: "Estante 2",
                Disponivel: false,
            },
        },
    }

    // Imprimir os detalhes do livro
    fmt.Println("Título:", livro.Titulo)
    fmt.Println("Autor:", livro.Autor)
    fmt.Println("Ano de Publicação:", livro.AnoPub)
    fmt.Println("ISBN:", livro.ISBN)


    livro.removerExemplar(2)

    // Imprimir os detalhes dos exemplares
    for _, exemplar := range livro.Exemplares {
        fmt.Printf("Exemplar %d: Localização=%s, Disponível=%t\n", exemplar.ID, exemplar.Localizacao, exemplar.Disponivel)
    }
}
