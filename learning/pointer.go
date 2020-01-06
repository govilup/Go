package main

import "fmt"

type Person struct {
	first string
	last  string
}

//create a method which will accept
//a pointer and change the value at
//at that address
func changeMe(p *Person) {
	//(*p.first) = "Bruce"
	//(*p.last) = "Wayne"
	//Above statements are same
	//as below
	p.first = "Bruce"
	p.last = "Wayne"
}

func main() {
	p := Person{
		first: "James",
		last:  "Bond",
	}
	fmt.Println("Before Calling changeMe function: ", p)
	//passing address to the function
	//which is of type pointer
	changeMe(&p)
	fmt.Println("After calling changeMe function: ", p)
}
