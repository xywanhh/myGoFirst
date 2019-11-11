package main

import (
	"fmt"
	"time"
)

/*
在一个 case 语句中，可以使用逗号来分隔多个表达式。

不带表达式的 switch 是实现 if/else 逻辑的另一种方式。这里展示了 case 表达式是如何使用非常量的。
*/
func main() {
	// a basic switch.
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")	
	}

	// You can use commas to separate multiple expressions in the same case statement. 
	// We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")	
	}

	// switch without an expression is an alternate way to express if/else logic. 
	// Here we also show how the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")	
	}

	// A type switch compares types instead of values. You can use this to discover the type of an interface value. 
	// In this example, the variable t will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)		
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")

}