package main

import "fmt"
import "math"

/*
Go supports constants of character,string,boolean and numberic values.

const declares a constant value.

A const statement can appear anywhere a var statement can.

常数表达式可以执行任意精度的运算

数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化。

当上下文需要时，一个数可以被给定一个类型，比如变量赋值或者函数调用。
*/

const s string = "hello"
func main()  {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(float64(d))

	// 举个例子，这里的 math.Sin函数需要一个 float64 的参数。
	fmt.Println(math.Sin(float64(n)))
	fmt.Println(math.Sin(n))

}