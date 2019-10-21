package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/* 验证euler公式 e^iπ + 1 = 0*/

func euler()  {
	c := 3 + 4i // 定义一个复数
	fmt.Println(cmplx.Abs(c)) // 求c的模 5

	fmt.Println(
		cmplx.Pow(math.E, 1i * math.Pi) + 1) // 求euler公式 e^iπ + 1 = 0 结果：(0+1.2246467991473515e-16i)
	fmt.Println(
		cmplx.Exp(1i * math.Pi) + 1) // 求euler公式 e^iπ + 1 = 0 结果：(0+1.2246467991473515e-16i)
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1) // 结果：(0.000+0.000i)

}

func main() {
	euler()
}
