package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
}

type secretAgent struct {
	Person
	ltk bool
}

//The below is a method
func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, "-the secret agent speak.")
}

func (p Person) speak() {
	fmt.Println("I am ", p.first, p.last, "-the person speak.")
}

func main() {
	sa1 := secretAgent{
		Person: Person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		Person: Person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: true,
	}

	p1 := Person{
		first: "Durga",
		last:  "Ojha",
	}

	//call method using the variable
	sa1.speak()
	sa2.speak()
	p1.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)
}

func bar(h human) {
	fmt.Println("bar is being called.", h)
}

//Below will create an interfce
//used for plolymorphism
type human interface {
	speak()
}
