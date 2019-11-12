package main

import "fmt"

/*
Functions are central in Go. 

Go 需要明确的返回值
Call a function just as you’d expect, with name(args).

多返回值
Multiple Return Values

允许可变长参数
Accepting a variable number of arguments
Variadic Functions.
fmt.Println is a common variadic function.

闭包
Closures.
Another key aspect of functions in Go is their ability to form closures
*/

func plus(a int, b int) int {
	return a + b
}

func pulsPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main()  {
	res := plus(1, 2)
	fmt.Println("1+2= ", res)

	res = pulsPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
	
	a, b := vals()
    fmt.Println(a)
	fmt.Println(b)
	
	_, c := vals()
	fmt.Println(c)

	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3}
	sum(nums...)
}