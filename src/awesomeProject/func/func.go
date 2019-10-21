package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		panic("unsupported operation: " + op)
	}
}
func eval1(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// 函数式编程改写
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args (%d, %d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += i
	}
	return s
}

func div(a, b int) (int, int)  {
	return a / b, a % b
}
func div1(a, b int) (q, r int)  {
	q, r = a / b, a % b
	return
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	//fmt.Println(eval(3, 4, "%"))
	fmt.Println(div(5, 2))
	fmt.Println(div1(5, 2))

	apply(pow, 3, 4)
	apply(func(i int, i2 int) int {
		return int(math.Pow(float64(i), float64(i2)))
	}, 3, 4)

	fmt.Println(sum(1, 2, 3, 4, 5))
}
