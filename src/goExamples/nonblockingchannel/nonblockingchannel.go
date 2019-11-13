
/*
非阻塞通道操作.

常规的通过通道发送和接收数据是阻塞的。
Basic sends and receives on channels are blocking.

然而，我们可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。
However, we can use select with a default clause to implement non-blocking sends, receives, 
and even non-blocking multi-way selects.

非阻塞发送
非阻塞接收

可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。

通道的关闭
Closing a channel indicates that no more values will be sent on it. 
This can be useful to communicate completion to the channel’s receivers.

通道遍历
一个非空的通道也是可以关闭的，但是通道中剩下的值仍然可以被接收到。
it’s possible to close a non-empty channel but still have the remaining values be received.
*/

package main
import "fmt"

func main() {
	m1 := make(chan string)
	s1 := make(chan bool)

	// Here’s a non-blocking receive
	select {
	case msg := <- m1:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")	
	}

	// Here’s a non-blocking send.
	msg := "hi"
	select {
	case m1 <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")	
	}

	select {
	case msg := <- m1:
		fmt.Println("received message", msg)
	case sig := <- s1:
		fmt.Println("received signal", sig)
	default:
        fmt.Println("no activity")		
	}

	//  the more value will be false if jobs has been closed and all values in the channel have already been received. 
	
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, ok := <-jobs
			if ok {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j:=0;j<=3;j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	q1 := make(chan string, 2)
	q1 <- "aa"
	q1 <- "bb"
	close(q1)

	for j := range q1 {
		fmt.Println(j)
	}

}