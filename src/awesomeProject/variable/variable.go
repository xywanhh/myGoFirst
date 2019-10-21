package main

import (
	"fmt"
	"math"
)

var aa = 123 // private 只能用在包内
// bb := 234 // 编译报错，短声明只能用在函数中
var AA = 123 // public 可以用在包外

// 使用var() 进行集中定义
var (
	AAA = 123
	BBB = 223
	SSS = "abc"
	Boo = true
)

func variableZeroValue()  {
	var a int
	var s string
	fmt.Println(a, s) // 变量定义有初始“零值”
	fmt.Printf("a:%d, s:%q\n", a, s) // 打印出字符串的“零值”
}

func variableInitialValue() {
	var a, b int = 3, 4 // 定义多个变量
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 1, 2, "abc", true // 类型推断
	var s = "def"
	fmt.Println(a, b, c, d, s)
}

func variableShorter() {
	a, b, c, d := 1, 2, "abc", true // 短声明，语法糖
	// b := 3 // 编译错误，变量不能重复定义
	s := "def"
	fmt.Println(a, b, c, d, s)
}

func triangle()  {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

// 常量值可以当作任意类型的值使用
const fileName = "a.txt" // 外部的常量

func consts() {
	const a = 3
	// const()
	const (
		b, c = 4, "abc"
	)
	var d int
	d = int(math.Sqrt(a*a + b*b))
	fmt.Println(d)
}

// 用常量来实现枚举
func enums()  {
	const (
		CPP = 0
		java = 1
		python = 2
		golang = 3
	)
	fmt.Println(CPP, java, python, golang)

	const (
		CPP1 = iota
		_ //
		java1
		python1
		golang1 = iota
		c
	)
	fmt.Println(CPP1, java1, python1, golang1, c)

	// b, kb, mb, gb, tb, pb
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(AA, BBB, aa, Boo) // 打印全局变量
	triangle()
	consts()
	enums()
}
