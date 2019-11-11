## 协程 Coroutine
```text
- 轻量集“线程”
- 非抢占式多任务处理，有协程主动交出控制全
- 编译器/解释器/虚拟机层面的多任务
- 多个协程可以在一个或者多个线程上运行


子程序是协程的一个特例

java不支持协程

goroutine的定义
- 任何函数只要机上go就能送个i调度器运行
- 不需要在定义是区分是否是一部函数
- 调度器在核实的点进行切换
- 使用-race来检测数据访问冲突

$ go run -race go1.go

goroutine肯恶搞的切换点：
- I/O, select
- channel
- 等待所
- 函数调用时（有事）
- runtime.Gosched()
以上只是参考，不能保证切换，不能保证在其他地方不切换
```

```text
【
goroutine <---- channel -----> goroutine
调度器
】
```

## channel
```text
channel
buffered channel
range
close
理论基础：communication sequential process (CSP)

不要通过共享内存来通信，通过通信来共享内存
```