package main

import "fmt"

// 函数式编程与闭包
// 函数式一等公民： 可以作为变量，参数，返回值使用

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// 正统的函数式编程写法
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}
