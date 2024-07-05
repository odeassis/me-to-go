package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID, ManagerID int
	Name, Address, Position string
	DoB time.Time
	Salary float64
}

type Address struct {
	Number int
	City string
	Cep string
}

type Person struct {
	Name string
	LastName string
	Age int
	Friends []Person
	Address
}

func newPerson() (newP Person) {
	newPerson := Person{
		Name: "",
		LastName: "",
		Age: 0,
		Friends: nil,
		Address: Address{},
	}

	return newPerson
}

func (p Person) show() {
	fmt.Printf("Nome: %s\n", p.Name)
	fmt.Printf("Last Name: %s\n", p.LastName)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Println("***************Address*****************")
	fmt.Printf("City: %s\n", p.City)
	fmt.Printf("CEP: %s\n", p.Cep)
	fmt.Printf("Number: %d\n", p.Number)
	fmt.Println("****************Friends****************")
	
	for i, v := range p.Friends {
		if len(v.Friends[i].Friends) == 0 {
			break
		}
		v.show()
	}
	fmt.Println("***************************************")
}

func (p *Person) addFriend(newFriend Person) {
	p.Friends = append(p.Friends, newFriend)
}

func main() {
	ana := Person{
		Name: "Anna",
		LastName: "Silva",
		Age: 23,
		Friends: []Person{},
		Address: Address{
			Number: 102,
			City: "Dom Pedro",
			Cep: "10101010",
		},
	}

	franc := newPerson()
	franc.Name = "o_deassis"

	ana.addFriend(franc)

	fmt.Println(franc)
	ana.show()

	var dilbert Employee

	dilbert.Salary -= 500
	position := &dilbert.Position

	*position = "Senior" + *position

	// var employeeOfTheMonth *Employee = &delbert
	employeeOfTheMonth := &dilbert

	employeeOfTheMonth.Position += " (proactive team player)"
}