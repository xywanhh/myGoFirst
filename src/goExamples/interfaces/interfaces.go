package main

import (
	"fmt"
	"math"
)
/*
Go’s mechanism for grouping and naming related sets of methods: interfaces.

Interfaces are named collections of method signatures.

To implement an interface in Go, we just need to implement all the methods in the interface. 
*/

// 定义接口
type geometry interface {
	area() float64
	perim() float64
}

// 结构体1
type rect struct {
	width, height float64
}
// 实现接口
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}


// 结构体2
type circle struct {
    radius float64
}
// 实现接口
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// 接收接口参数
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main()  {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	
	measure(r)
    measure(c)
}