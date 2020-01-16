package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Global variable whose scope is entire package
var wg sync.WaitGroup

func main() {
	fmt.Println("GOOS\t\t", runtime.GOOS)
	fmt.Println("GOARACH\t\t", runtime.GOARCH)
	fmt.Println("NumCPU\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	fmt.Println()

	wg.Add(1)
	//add go to start a new go routine.
	//only adding go foo() will not run foo()
	//the moment next go routine start the code
	//of func main is executed completly
	//To avoid that we will use synchronization
	//using primitives Wait group functions.
	go foo()
	bar()

	fmt.Println("NumCPU\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	fmt.Println("Foo is Printing")
	fmt.Println("---------------")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func bar() {
	fmt.Println("Bar is Printing")
	fmt.Println("---------------")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
