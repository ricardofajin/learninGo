package main

import "fmt"

//This function intSeq returns another function, which we define anonymously in the body of intSeq.
func intSec() func() int {
	i := 0
	return func() int {
		i++
		return i //The returned function closes over the variable i to form a closure.
	}
}

func main() {
	//We call intSeq, assigning the result (a function) to nextInt.
	nextInt := intSec() //This function value captures its own i value, which will be updated each time we call nextInt.

	//See the effect of the closure by calling nextInt a few times.
	println("First println -> ", nextInt())
	println("Second println -> ", nextInt())
	println("Third println -> ", nextInt())

	//To confirm that the state is unique to that particular function, create and test a new one.
	nextInts := intSec()
	fmt.Println("Re-declaring the variable -> ", nextInts())

}
