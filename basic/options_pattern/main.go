package main

import (
	"errors"
	"fmt"
)

type Customer struct {
	id        string
	email     string
	firstName *string // ponteiro para `string`, pode ser nulo (ou seja, não preenchido)
	lastName  *string
	points    int
}

var ErrEmptyId = errors.New("Id cannot be empty")
var ErrInvalidEmail = errors.New("email is invalid")
var ErrEmptyFirstName = errors.New("first name cannot be empty")
var ErrEmptyLastName = errors.New("last name cannot be empty")

func (c *Customer) validate() error {
	if c.id == "" {
		return ErrEmptyId
	}
	if c.email == "" {
		return ErrInvalidEmail
	}
	if c.firstName != nil && *c.firstName == "" {
		return ErrEmptyFirstName
	}
	if c.lastName != nil && *c.lastName == "" {
		return ErrEmptyLastName
	}
	return nil
}

type Option func(c *Customer)

/*
A `Option` que a função `WithName` retorna é, na verdade, uma função anônima (closure)
que, quando chamada, atribui os valores dos ponteiros de string `firstName` e `lastName`
aos campos correspondentes da estrutura `Customer`.
*/
func WithName(firstName, lastName *string) Option {
	return func(c *Customer) {
		c.firstName = firstName
		c.lastName = lastName
	}
}

func WithPoint(points int) Option {
	return func(c *Customer) {
		c.points = points
	}
}

/*
`NewCustomer` é o contrutor do tipo `Customer`. Ela cria uma nova instância de `Customer`
com `id` e `email` fornecidos. Além disso, ela aceita um número variável de argumentos do tipo
`Option`. Isso significa que, ao criar um novo cliente, pode-se fornercer várias opções que
ajustam o cliente de acordo com suas necessidades.
*/
func NewCustomer(id, email string, option ...Option) (*Customer, error) {
	c := &Customer{
		id:     id,
		email:  email,
		points: 100,
	}

	/*
		Ela itera sobre todas as opções fornecidas, chamando cada uma delas com o clinete `c`
		como argumento. Isso permite que as opções modifiquem o cliente conforme necessário.
	*/
	for _, opt := range option {
		opt(c)
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func main() {
	id := "6fa49e0a"
	email := "jane@doe.com"
	firstName := "Jane"
	lastName := "Doe"
	_, _ = NewCustomer(id, email)
	_, _ = NewCustomer(id, email, WithName(&firstName, &lastName))
	_, _ = NewCustomer(id, email, WithPoint(200))
	customer, _ := NewCustomer(id, email, WithName(&firstName, &lastName), WithPoint(200))

	fmt.Println(*customer.firstName)
}
