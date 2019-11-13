package main

	
import "time"
import "fmt"

/*
超时处理
得益于通道和 select，在 Go中实现超时操作是简洁而优雅

Note that the channel is buffered, so the send in the goroutine is nonblocking. 
*/
func main()  {
	
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")	
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")	
	}




}