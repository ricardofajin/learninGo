package main

/*
	Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.
*/

import "fmt"

//This person struct type has name and age fields.
type person struct {
	name string
	age  int
}

func main() {

	//This syntax creates a new struct.
	fmt.Println("Syntax to initialize the struct - person{'name_val',age_val} -> ", person{"Bob", 20})

	//You can name the fields when initializing a struct.
	fmt.Println("Using name fields when initialize - person{name:'name_val',age:age_val} -> ", person{name: "Alice", age: 30})

	//Omitted fields will be zero-valued.
	fmt.Println("Using name fields and dont use the age field when initialize - person{name:'name_val'} -> ", person{name: "Fred"})

	//An & prefix yields a pointer to the struct.
	fmt.Println("Using (&)yield a pointer to the structure - &person{name:'name_val'} -> ", &person{name: "Ann", age: 50})

	//Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println("Access struct fields with a dot. - s.name -> ", s.name)

	//You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println("use dots with struct pointers - sp.age -> ", sp.age)

	//Structs are mutable.
	sp.age = 51
	fmt.Println("Structs are mutable sp.age=51 - sp.age -> ", sp.age)
}
