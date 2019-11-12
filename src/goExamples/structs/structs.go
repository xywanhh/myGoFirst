package main

import "fmt"

/*
Go’s structs are typed collections of fields.

Access struct fields with a dot.

结构体是可变的。
Structs are mutable.
*/
type person struct {
	name string
	age int
}

func NewPerson(name string) *person {
	p := person{name : name}
	p.age = 1
	return &p
}

func main()  {
	fmt.Println(person{"alic", 20})

	fmt.Println(person{name:"alic", age:11})

	fmt.Println(person{name:"AAA"})

	fmt.Println(&person{name:"ANNN", age:222})

	fmt.Println(NewPerson("JHon"))

	s := person{name:"SS", age:333}
	fmt.Println(s.name)

	// 对结构体指针使用. - 指针会被自动解引用。
	s1 := &s
	fmt.Println(s1.age)

	s1.name = "SS1"
	fmt.Println(s.name)

}