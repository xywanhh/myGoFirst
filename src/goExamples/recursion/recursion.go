package main

import "fmt"

/*
递归
Recursion

临界点
数学归纳法
*/

// 求阶乘
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func main()  {
	fmt.Println(fact(6))
}