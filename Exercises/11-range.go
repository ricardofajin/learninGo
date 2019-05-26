package main

import "fmt"

/*

Range iterates over elements in a variety of data structures.
Let’s see how to use range with some of the data structures we’ve already learned.

*/

func main() {

	//Here we use range to sum the numbers in a slice. Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { //Here we didn’t need the value itself, so we ignored it with the blank identifier _.
		sum += num
	}
	fmt.Println("Print variable sum -> ", sum)

	//range on arrays and slices provides both the index and value for each entry.
	//Above we didn’t need the index, so we ignored it with the blank identifier _.
	for i, num := range nums { //Sometimes we actually want the indexes though.
		if num == 3 {
			fmt.Println("index with number 3, Position ->", i)
		}
	}

	//range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) //Printf read %s %s
	}

	//range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("Key", k)
	}

	//range on strings iterates over Unicode code points.
	for i, c := range "go" { //The first value is the starting byte index of the rune and the second the rune itself.
		fmt.Println(i, c)
	}
}
