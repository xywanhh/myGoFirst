package main

import (
	"fmt"
	"config"
)

func main() {
	config.LoadConfig()
	fmt.Println("hello xinxin")
}
