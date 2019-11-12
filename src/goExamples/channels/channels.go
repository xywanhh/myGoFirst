package main

import (
	"fmt"
	"time"
)

/*
Channels are the pipes that connect concurrent goroutines. 
You can send values into channels from one goroutine and receive those values into another goroutine.

Create a new channel with make(chan val-type). Channels are typed by the values they convey.

Send a value into a channel using the channel <- syntax.
The <-channel syntax receives a value from the channel.

默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。
By default sends and receives block until both the sender and receiver are ready. 
This property allowed us to wait at the end of our program for the "ping" message 
without having to use any other synchronization.

通道缓冲
默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan）通道准备好接收时，才允许进行发送（chan <-）。
可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值。

通道同步
Channel Synchronization
We can use channels to synchronize execution across goroutines. 

通道方向
当使用通道作为函数的参数时，你可以指定这个通道是不是只用来发送或者接收值。这个特性提升了程序的类型安全性。
ping 函数定义了一个只允许发送数据的通道。尝试使用这个通道来接收数据将会得到一个编译时错误。
pong 函数允许通道（pings）来接收数据，另一通道（pongs）来发送数据。

*/

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

// ping 函数定义了一个只允许发送数据的通道。尝试使用这个通道来接收数据将会得到一个编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}
// pong 函数允许通道（pings）来接收数据，另一通道（pongs）来发送数据。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}


func main()  {
	messages := make(chan string)
	
	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)


	m1 := make(chan string, 2)
	// 因为这个通道是有缓冲区的，即使没有一个对应的并发接收方，我们仍然可以发送这些值。
	m1 <- "aa"
	m1 <- "bb"

	fmt.Println(<-m1)
	fmt.Println(<-m1)

	// using a blocking receive to wait for a goroutine to finish.
	// When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
	// The done channel will be used to notify another goroutine that this function’s work is done.
	done := make(chan bool)
	go worker(done)
	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	// 如果把 <- done 这行代码从程序中移除，程序甚至会在 worker还没开始运行时就结束了。
	<-done
	
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passing message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	
}