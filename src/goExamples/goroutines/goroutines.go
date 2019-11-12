package main

import (
	"fmt"
	"time"
)

/*
A goroutine is a lightweight thread of execution.

Go 协程的辅助特性：通道。
*/

func f(from string) {
	for i:=0;i<3;i++ {
		fmt.Println(from, " : ", i)
	}
}

func main()  {
	// running it synchronously.
	f("direct")	

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

}