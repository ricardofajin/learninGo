package Exercises

import "fmt"

func main() {

	//var declares 1 or more variables.
	var a = "string variable"
	fmt.Println(a)
	//An exemple with declaration of 2 differents variables but the same attribute(integer)
	var b, c int = 1, 2
	fmt.Println(b,c)

	//Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	//Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	//The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short" in this case.
	f := "String 'f' initialized with :="
	fmt.Println(f)

}
