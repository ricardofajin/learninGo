package main

/*
	Functions are central in Go.
	We’ll learn about functions with a few different examples.
*/

import "fmt"

//Here’s a function that takes two ints and returns their sum as an int.

func plus(a int, b int) int {

	return a + b
}

//When you have multiple consecutive parameters of the same type,
//you may omit the type name for the like-typed parameters up to the final parameter that declares the type.

func plusPlus(a, b, c int) int {

	return a + b + c
}

func main() {

	//Call a function just as you’d expect, with name(args)
	res := plus(1, 2)
	fmt.Println("Output furst function, plus (1 + 2) -> ", res)

	res2 := plusPlus(1, 2, 3)
	fmt.Println("Output furst function, plusPlus (1 + 2 + 3) -> ", res2)
}
