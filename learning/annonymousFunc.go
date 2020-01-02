package main

import (
	"fmt"
)

func main() {
	foo() //this is not an annonymous function

	//annonymous functions are functions
	//without any name and are called
	//inside some function
	func() {
		fmt.Println("Annonymous function is being called.")
	}() //these are required to pass any argument

	func(x int) {
		fmt.Println("Annonymous function with value", x, "is called")
	}(38)

}

func foo() {
	fmt.Println("foo ran")
}
