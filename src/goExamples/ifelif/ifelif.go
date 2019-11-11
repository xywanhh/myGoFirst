package main

import "fmt"

/*
没有三元运算符

*/
func main() {
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8 % 4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals; any variables declared in this statement are available in all branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 0 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}