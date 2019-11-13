package main

import (
    "fmt"
    "time"
)

/*
Go的通道选择器 让你可以同时等待多个通道操作。
Go协程和通道以及选择器的结合是 Go 的一个强大特性。

*/
func main()  {
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0;i < 2;i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)	
		}
	}


}