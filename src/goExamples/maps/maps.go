package main

import "fmt"

/*
关联数据类型

make(map[key-type]val-type).


*/
func main()  {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)
	
	// name[key] = val
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// len returns the number of key/value pairs
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map: ", m)

	// 当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。这可以用来消除键不存在和键有零值
	_, ok := m["k2"]
	fmt.Println(ok)

	// declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("mapN: ", n)

}