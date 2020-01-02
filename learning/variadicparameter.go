package main

import "fmt"

func main() {
	//Using short declaration operator
	xs := []int{2, 4, 6, 8, 10} //slice of int, composite literal

	//Call a function which accepts variadic parameter
	//three dots after the variable name will unfurl
	//the values inside the slice
	x := foo(xs...)
	fmt.Println(x)

}

//three dots before the type of the function
//will accepts multiples values
//This is called variadic parameter
func foo(x ...int) int {
	fmt.Println(x)
	//print the type of x
	fmt.Printf("%T\n", x)

	//Print the sum
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
