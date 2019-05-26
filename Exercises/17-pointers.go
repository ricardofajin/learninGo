package main

/*
	Go supports pointers, allowing you to pass references to values and records within your program.
	- https://en.wikipedia.org/wiki/Pointer_(computer_programming)
*/
import "fmt"

//We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr.
func zeroval(ival int) { //zeroval has an int parameter, so arguments will be passed to it by value.
	ival = 0 //zeroval will get a copy of ival distinct from the one in the calling function.
}

//zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
func zeroptr(iptr *int) {
	*iptr = 0 //The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.
	//Assigning a value to a dereferenced pointer changes the value at the referenced address.
}

func main() {
	i := 1
	fmt.Println("initial -> ", i)

	zeroval(i)
	fmt.Println("Using zeroval(i) - Println(i) -> ", i)

	zeroptr(&i) //The &i syntax gives the memory address of i, i.e. a pointer to i.
	fmt.Println("Using zeroptr(&i) - Println(i) -> ", i)

	fmt.Println("Using only Println(&i) to see the pointer used -> ", &i) //Pointers can be printed too.

	/*
		zeroval doesn’t change the i in main, but zeroptr does
		because it has a reference to the memory address for that variable.
	*/

}
