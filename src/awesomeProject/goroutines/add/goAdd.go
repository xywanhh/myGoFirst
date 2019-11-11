package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i:=0;i<10;i++ {
		go func(i int) {
			for {
				a[i]++
				// 这里的协程可能会一直占用，不交出控制权，导致一直卡主
				// runtime.Gosched() 手动交出控制权
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
