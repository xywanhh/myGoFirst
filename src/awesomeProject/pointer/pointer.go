package main

import "fmt"

// 指针不能做运算
// Go语言只有值传递一种方式

func swap(a, b int)  {
	a, b = b, a
}
func swapEffect(a, b *int)  {
	*a, *b = *b, *a
}

func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
	fmt.Println(*pa)

	a, b := 1, 2
	swap(a, b)
	fmt.Println(a, b)

	swapEffect(&a, &b)
	fmt.Println(a, b)
}
