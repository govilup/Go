package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{4, 3, 6, 18, 20, 1, 5}

	fmt.Println("Before Sort :", numbers)
	sort.Ints(numbers)
	fmt.Println("After Sort :", numbers)

	//custom sort
	p1 := Person{"Bruce", "Wayne", 28}
	p2 := Person{"James", "Bond", 32}
	p3 := Person{"Logan", "", 200}

	people := []Person{p2, p1, p3}

	fmt.Println(people)
	// There are two ways to sort a slice. First, one can define
	// a set of methods for the slice type, as with ByAge, and
	// call sort.Sort. In this first example we use that technique.
	sort.Sort(ByAge(people))
	fmt.Println(people)
}

//Custom sort
type Person struct {
	First string
	Last  string
	Age   int
}

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByAge []Person
