package main

import (
	"fmt"
)

//One scope enclosing other scopes:
//Variables declared in the outer scope
//are accessible in inner scopes
//Closure helps us in limits the scopes of a variable.

func main() {
	//Variable a and b is of type func
	a := incrementor()
	b := incrementor()
	fmt.Println("Incremented value of a ", a())
	fmt.Println("Incremented value of b ", b())
	fmt.Println("Incremented value of a ", a())
	fmt.Println("Incremented value of a ", a())
}

func incrementor() func() int {
	//the scope of the variable
	//x is this complere block
	//of function
	var x int
	//Do we need to pass te variable x
	//in the return func() ?
	//No, since scope of x is
	//entire incrementor() block.
	return func() int {
		x++
		return x
	}
}
