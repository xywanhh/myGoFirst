/*
定时器表示在未来某一时刻的独立事件。你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。

<-timer1.C 直到这个定时器的通道 C 明确的发送了定时器失效的值之前，将一直阻塞。
The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer expired.

如果你需要的仅仅是单纯的等待，你需要使用 time.Sleep。定时器是有用原因之一就是你可以在定时器失效之前，取消这个定时器。

打点器
Tickers
ticks periodically until we stop it.
Tickers use a similar mechanism to timers: a channel that is sent values.
Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more values on its channel.
*/

package main
import (
    "fmt"
    "time"
)
func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
        fmt.Println("Timer 2 stopped")
	}
	
	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at: ", t)	
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}