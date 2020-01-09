package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	//To Marshal, start the name of
	//variable with Capital letter
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := Person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []Person{p1, p2}

	//Marshal person into json

	ps, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ps))
}
