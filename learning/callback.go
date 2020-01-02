package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := sum(ii...)
	fmt.Println("Sum of the numbers in the slice is ", s)

	//Call a callback func
	s2 := even(sum, ii...)
	fmt.Println("Sum of even numbers is ", s2)

	s3 := odd(sum, ii...)
	fmt.Println("Sum of odd numbers is ", s3)
}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//Callback is when we pass func as an argument.
func even(f func(xi ...int) int, ii ...int) int {
	var yi []int
	for _, v := range ii {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	//Unfurl the slice
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}
