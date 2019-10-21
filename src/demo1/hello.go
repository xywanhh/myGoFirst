package main

import "fmt"

func SayHello(s string) string {
	return "hello " + s 
}

func main() {
	fmt.Println(SayHello("baobao"))
}