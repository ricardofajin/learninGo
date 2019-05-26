package main

import "fmt"

/*
	Go supports methods defined on struct types.
*/

type rect struct {
	widith, height int
}

//This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.widith * r.height
}

//Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
func (r rect) perim() int {
	return 2*r.widith + 2*r.height

}

func main() {

	r := rect{widith: 10, height: 5}

	//Here we call the 2 methods defined for our struct.
	fmt.Println("Calculate area r.area()-> ", r.area())
	fmt.Println("Calculate area r.perim()-> ", r.perim())

	//Go automatically handles conversion between values and pointers for method calls.
	rp := &r //You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.

	fmt.Println("Calculate area r.area()-> ", rp.area())
	fmt.Println("Calculate area r.perim()-> ", rp.perim())

}
