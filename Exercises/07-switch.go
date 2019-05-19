package Exercises

import "fmt"
import "time"

// Switch statements express conditionals across many branches.

func main() {
	//Here’s a basic switch.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}


	//You can use commas to separate multiple expressions in the same case statement.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	//We use the optional default case in this example as well.
	default:
		fmt.Println("It's a weekday")
	}

	//switch without an expression is an alternate way to express if/else logic.
	t := time.Now()

	//Here we also show how the case expressions can be non-constants.
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//A type switch compares types instead of values. You can use this to discover the type of an interface value.
	//In this example, the variable t will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a integer")
		default:
			fmt.Printf("I Don't know my type, but Printf can help: %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey you")
}
