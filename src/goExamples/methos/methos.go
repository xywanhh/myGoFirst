package main

import "fmt"

/*
Go supports methods defined on struct types.

Methods can be defined for either pointer or value receiver types. 
*/

type rect struct {
	width, height int
}

// 指针接收者
func (r *rect) area() int {
	return r.width * r.height
}

// 值接收者
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main()  {
	r := rect{width:10, height:5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	
	rp := &r
	fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}