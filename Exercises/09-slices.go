package Exercises

import (
	"fmt"
)

//You’ll see slices much more often than arrays in typical Go.
//Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

func main() {
	//Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	s := make([]string, 3) //To create an empty slice with non-zero length, use the builtin make.
	fmt.Println("emp: ", s) //Here we make a slice of strings of length 3 (initially zero-valued).

	//We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	//len returns the length of the slice as expected.
	fmt.Println("len:", len(s))


	//In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append

	s = append(s, "d") //which returns a slice containing one or more new values.
	s = append(s, "e", "f") // Note that we need to accept a return value from append as we may get a new slice value.
	fmt.Println("Append: ", s)

	//Slices can also be copy’d. Here we create an empty slice c of the same length as s
	c := make([]string, len(s))
	//and copy into c from s.
	copy(c, s)
	fmt.Println("Copy:", c)

	//Slices support a “slice” operator with the syntax slice[low:high].
	l := s[2:5] // For example, this gets a slice of the elements s[2], s[3], and s[4].
	fmt.Println("slice 1 -  this gets a slice of the elements s[2], s[3], and s[4]:\n", l)

	l = s[:5] //This slices up to (but excluding) s[5].
	fmt.Println("slice 2 - This slices up to (but excluding) s[5]:\n", l)

	l = s[2:] //And this slices up from (and including) s[2]
	fmt.Println("slice 3 - And this slices up from (and including) s[2]:\n", l)


	//We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("declare and initialize:\n", t)


	//Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary,
	// unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2 dimentions using slice:\n", twoD) //Note that while slices are different types than arrays, they are rendered similarly by fmt.Println.
}
