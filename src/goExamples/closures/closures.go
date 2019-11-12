package main

import "fmt"

/*
Go supports anonymous functions, which can form closures.
*/

// 这个 intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数。这个返回的函数使用闭包的方式 隐藏 变量 i。

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main()  {
	
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 这个状态对于这个特定的函数是唯一的
	newInts := intSeq()
	fmt.Println(newInts())
}