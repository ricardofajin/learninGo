package main

/*

	Maps are Go’s built-in associative data type (https://en.wikipedia.org/wiki/Associative_array)
   (sometimes called hashes or dicts in other languages).

*/

import "fmt"

func main() {
	//To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	//Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	//Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("map: ", m)

	//Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("Variable 1(v1) receive k1 -> ", v1)

	//The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("Number of the key/value - len:", len(m))

	//The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map after deleted key(2) -> ", m)

	//The optional second return value when getting a value from a map indicates if the key was present in the map.
	//This can be used to disambiguate between missing keys and keys with zero values like 0 or "".

	_, prs := m["k2"] //Here we did not need the value itself, so we ignored it with the blank identifier _.
	fmt.Println("prs: ", prs)

	//You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"afoo": 1, "bar": 2}

	//the maps initialize the map with alphabetical order. (new feature?)
	fmt.Println("new map, initialized at the same line -> ", n) //Note that maps appear in the form map[k:v k:v] when printed with fmt.Println.

	fmt.Printf("+%v\n", n)

	for k, v := range n {
		fmt.Println("key: ", k)
		fmt.Println("value: ", v)
	}

	//? - In my case, the "bar" was started at the front of the "foo" -? why o.O
}
