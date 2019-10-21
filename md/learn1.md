go run
go build

Golang是静态类型语言，js是动态类型语言

Go负责内存分配和垃圾回收

推荐使用
int
float64

检查变量类型
reflect包
reflect.TypeOf(varbarie)

import myfmt "fmt" // 使用别名
import . "fmt" // 省略名称


type (
    newType int
    type1 float32
    type2 string
    type3 byte
)

// 常量组的定义
const (
    PI = 3.14
    const1 = "1"
    const2 = 1
)

// 全局变量组的声明，不能用在函数里
var (
    name = "abc"
    name1 = 1
)

Go基本类型

布尔bool
false/true

整型 int/uint

8位整型 int8/uint8 
长度1字节

字节型 byte(uint8别名)


16位整型 int16/uint16 
长度2字节

32位整型 int32(rune)/uint32 
长度4字节


64位整型 int64/uint64
长度8字节

浮点型 float32/ float64
长度4/8字节

值类型
array/struct/string

引用类型
slice/map/chan

接口类型  interface
函数类型 func

全局变量的声明可以使用var()组来声明
全局变量的声明不可以省略var，但可以并行声明
局部变量的声明不可以用var()组，只能并行声明


类型转换
strconv包
```go
var s string = "abc"
b, err = strconv.ParseBool(s)

s := strconv.FormatBool(true)
ftm.Println(s)

var a int = 65
b := strconv.Itoa(a) // 65
a, _ = strconv.Atoi(b) // 65

```

Go是静态类型语言，声明变量时必须显式或者隐式的指明变量类型
变量声明后不能重复声明
禁止将变量设置为nil，将导致编译错误

```go
var s string = "abc"
var sss = "bcd"

var ss string
ss = "abcd"
```

快捷变量声明
```go
var a,b,c string = "a","b","c"

var (
    s string = "abc"
    f bool = true
    i int
)
```

检查变量为空
```go
var s string
if s == "" {
    ftm.Printf("s has not been assigned a value and is zero value")
}
```

简短变量声明
只能在函数中使用
```go
s := "abc"
```

指针
*
&

常量
常量声明后，只能使用，不能修改
```go
const greeting string = "hello"
```

函数在Go语言里作为一种类型，一等公民，因为可以作为参数进行传递

函数签名
```go
func （int,int) int

func add(x, y int) int {
    return x + y
}
```

返回多个值
```go
func getGood() (int, string) {
    i := 1
    s := "abc"
    return i, s
}
```

不定参数
```go
...
func add(number... int) int {}

func add(numbers... int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    result := add(1,2,3,4)
}
```

具名返回值
```go
func say() (x string, y string) {
    x := "hello"
    y := "baobao"
    return
}
```

函数作为值传递
```go
func anotherFunc(f func() string) string {
    return f()
}

func main() {
    f := func() string {
        return "function called"
    }
    fmt.Println(anotherFunc(f))
}
```

逻辑控制
if,else, else if
switch
for
defer
goto


defer
在函数返回前执行一个函数
```go
func main() {
    defer fmt.Println("I am run after the function complete")
    ftm.Println("hello")
}
```

数组
```go
var sArr [3]string
sArr[0] = "a"
```

切片
```go
var aSlice = make([]string, 2)
aSlice[0] = "a"

aSlice = append(aSlice, "b")
// 使用append来达到删除的目的
aSlice = append(aSlice[:2], aSlice[2+1:])
copy(aSlice1, aSlice2) 
```

映射
```go
var aMap := make(map[string]int)
aMap["a"] = 1

delete(aMap,"a")
```

结构体
是由元素组成的结构
复杂的数据结构
模板

结构体并非创建面向对象的方式，是一种数据结构创建方式，旨在满足数据建模要求。

```go
type Movie struct {
    Name string
    Rating float64
}

func main() {
    m := Movie {
        Name : "abc",
        Rating: 1.1,
    }

    var mm Movie // 创建一个结构体的实例
    mm.Name = "abc"

    mmm := new(Movie) // 通过new创建结构体实例
    mmm.Name = "abc"

    am := Movie(Name:"abc", Rating:1.1)
    bm := Movie("abc", 1.1)
    c := Movie {
        Name : "abc",
        Rating: 1.1,
    }

    fmt.Println(m.Name, m.Rating)
}
```

结构体嵌套
```go
type Address struct {
    Number int
    Street string
    City string
}

type SuperHero struct {
    Mame string
    Age int
    Address Address
}

func main() {
    ironMan := SuperHero {
        Name : "ironman",
        Age : 36,
        Address : {
            Number : 123,
            Street : "Mountation",
            City : "Goman",
        },
    }

    fmt.Printf("%v\n", ironman)
}
```

Go语言中的零值
布尔型  bool    false
整型    int     0
浮点    float   0.0
字符串  string  ""
指针    pointer nil
函数    func    nil
接口    interface   nil
切片    slice   nil
通道    channel nil
映射    Map     nil

结构体的比较 == != 
类型和值相比
不能比较类型不同的结构体，会导致编译报错

检查结构体的类型
reflect.TypeOf(ironman)

结构体中首字母大写的可被导出，否则不导出

使用结构体时，明确指针引用和值引用
batman := ironman // 结构体的副本
batman := &ironman // 指向原始结构体的指针

接口
描述方法集的方式
实现模块化的强大方式
方法集的蓝本

方法
方法是一个接受被称为接收者的特殊参数的函数
接收者可以是指针，也可以是值

给结构体添加方法
```go
type Movie struct {
    Name string
    Rating float32
}

// 方法
func (m *Movie) summary() string {
    r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
    return m.Name + " , " + r
}

// 函数
func summary(m *Movie) string {
    r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
    return m.Name + " , " + r
}

func main() {
    m := Movie {
        Name : "ironman",
        Rating : 1.1,
    }
    fmt.Println(m.summary())
}
```

方法集
方法集是对特定数据类型进行调用的一组方法。
能够在数据类型和方法之间建立关系。
是一种封装功能的有效方式

接口以声明的方式提供了多态

```go
package main

import (
    "fmt"
    "errors"
)
// 定义接口
type Robot interface {
    PowerOn() error
}

// 实现接口1
type T21 struct {
    Name string
}

func (a *T21) PowerOn() error {
    return nil
}

// 实现接口2
type T22 struct {
    Broken bool
}

func (r *T22) PowerOn() error {
    if r.Broken {
        return errors.New("T22 is broken")
    } else {
        return nil
    }
}

// 接口也是一种数据类型，可以作为参数传递
func Boot(r Robot) error {
    return r.PowerOn()
}
```

字符串
在Go中，字符串是不可变的

字符串为字节切片

strings包
strings.ToLower("ABC")
strings.Index("abcd", "b")
strings.TrimSpace(" abc ")
创建字符串字面量
s := "abc"

rune字面量
\n
\t
\\
\'
\\"

只能是字符串类型进行拼接，整型和字符串拼接会报编译错误

可以使用strconv包的Itoa方法进行转换
```go
var i int = 1
var s string = "abc"
ii := strconv.Itoa(i)
var ss string = s + ii
```

使用缓冲区来进行字符串拼接
```go
package main

import "fmt"

func main() {
    var buffer byte.Buffer
    for i:=0; i <500; i++ {
        buffer.WriteString("z")
    }
    fmt.Println(buffer.String())
}
```

计算机底层实际上是将字符存储为数字的。

Unicode标准
UTF-8实现了Unicode标准


错误
错误作为类型，可以作为参数进行传递给函数或者方法

```go
package main

import (
    "fmt"
    "io/ioutil"
)

// func ReadFile(filename string) ([]byte, error)
func main() {
    file, err := ioutil.ReadFile("f.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("%s", file)
}
```

错误是一个值
标准库声明了接口error
Go独特的错误处理方式-由调用者处理错误
```go
type error interface {
    Error() string
}
```

标准库提供errors包用来创建和操作错误
```go
err := errors.New("abc")
if err != nil {
    fmt.Println(err)
}
```

设置错误的格式
标准库fmt包提供了Errorf，用来设置错误字符串的格式
```go
name, role := "a","b"
err := fmt.Errorf("The error %v %v quit", name, role)
```

eg:
```go
package main

import "fmt"

func Half(number int) (int, error) {
    if number % 2 != 0 {
        return -1, fmt.Errorf("cannot half %v", number)
    }
    return number%2, nil
}

func main() {
    n, err := Half(19)
    // 错误不是在函数中进行处理，而是在调用处进行处理，提供了极大的灵活性
    if err != nil {
        fmt.Println(n)
        return
    }
    fmt.Println(n)
}
```

慎用panic
panic是Go中的一个内置函数，终止正常的程序流程，导致程序停止运行，没有回旋余地

```go
func main() {
    fmt.Println("this is executed")
    panic("oh no,i can do no more ,goodbye")
    fmt.Println("this is not executed")
}
```

使用panic的场景
- 程序发生无法恢复的错误，再往下走会带来更多的问题
- 发生了无法处理的错误

并发和并行
并发就是同时处理很多事情。
并行就是同时做很多事情。

time.Sleep(time.Second * 2)
```go
package main

import (
    "fmt"
    "time"
)

func slowFunc() {
    time.Sleep(time.Second * 2)
    fmt.Println("func 1 finished")
}

func main() {
    slowFunc()
    fmt.Println("finally finished")
}
```

```go
package main

import (
    "fmt"
    "time"
)

func slowFunc() {
    time.Sleep(time.Second * 2)
    fmt.Println("func 1 finished")
}

func main() {
    go slowFunc() // 调用后会立即返回，执行后续的代码，来不及等待结果返回
    fmt.Println("finally finished")
}
```

```go
package main

import (
    "fmt"
    "time"
)

func slowFunc() {
    time.Sleep(time.Second * 2)
    fmt.Println("func 1 finished")
}

func main() {
    go slowFunc() // 调用后会立即返回，执行后续的代码，来不及等待结果返回
    fmt.Println("finally finished")
    time.Sleep(time.Second * 3) // 等待结果返回
}
```

Rob Pike "Concurrency Is Not Parallelism"


通道
通道使得Goroutine之间进行通信

通道与Goroutine互为补充，它让Goroutine能够互相通信。

《Effective Go》
“不同通过共享内存来通信，而是通过通信来共享内存”

通道的创建语法：
c := make(chan string)

向通道发送消息
c <- "a"

从通道里接受消息
msg := <-c

并发例子改进
```go
package main

import (
    "fmt"
    "time"
)

func slowFunc(c chan string) {
    time.Sleep(time.Second * 2)
    c <- "func 1 finished"
}

func main() {
    c := make(chan string)
    go slowFunc(c)

    msg := <-c
    fmt.Println(msg)
}
```

缓冲通道
c := make(chan string, 2) // 缓冲区长度
超过长度，导致错误

在知道需要启动多少个Goroutine或者需要限制调度的工作量时，缓冲通道很有用
```go
package main

import (
    "fmt"
    "time"
)

func receive(c chan string) {
    for msg := range c {
        fmt.Println(msg)
    }
}

func main() {
    c := make(chan string, 2) // 创建缓冲区长度为2的通道，最多能缓冲2条消息
    c <- "a" // 向通道发送消息
    c <- "b"

    Close(c) // 关闭通道，禁止再向通道发送消息

    recevie(c)
}
```

指定通道是只读的
c <-chan string

指定通道是只写的
c chan<- string

制定通道是可读写的
c chan string

```go
func channelReader(c <-chan string) {
    msg := <-c
    fmt.Println(msg)
}

func channelWriter(c chan<- string) {
    c <- "a"
}

func channelReaderAndWriter(c chan string) {
    msg := <-c
    fmt.Println(msg)
    c <- "a"
}
```

select
假如有多个Goroutine，执行最先返回的Goroutine

```go
c1 := make(chan string)
c2 := make(chan string)

select {
    case msg1 := <-c1:
        fmt.Println(msg1)
    case msg2 := <-c2:
        fmt.Println(msg2)
}

最先从哪个通道接受到消息，就执行哪个分支
```

完整的例子
```go
package main

import (
    "fmt"
    "time"
)

func ping1(c chan string) {
    time.Sleep(time.Second * 1)
    c <- "ping 1 finished"
}

func ping2(c chan string) {
    time.Sleep(time.Second * 2)
    c <- "ping 2 finished"
}

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go ping1(c1)
    go ping2(c2)

    select {
        case msg1 := <-c1:
            fmt.Println(msg1)
        case msg2 := <-c2:
            fmt.Println(msg2)
        case <-time.After(500 * time.Millisecond): // 超时500ms，执行该分支
            fmt.Println("no message received in 500 mills")        
    }
}

// ping1 finished
```

使用推出通道
```go
package main

import (
    "fmt"
    "time"
)

func sender(c chan string) {
    t := time.NewTicker(1 * time.Second)
    c <- "sending message"
    <-t.C
}

func main() {
    c := make(chan string)
    stop := make(chan bool) // 创建推出通道

    go sender(c)

    go func() {
        time.Sleep(time.Second * 2)
        fmt.Println("Time's up")
        stop <- true
    }()

    for {
        select {
            case <-stop:
                return
            case msg := <-c:
                fmt.Println(msg)
        }
    }
}

// sending message
// sending message
// Time's up
```


使用包实现代码重用

安装第三方包
go get 第三方包在远程服务器的路径

go get github.com/golang/example/stringutil
安装到GOPATH的路径下
%GOPATH%\src\github\golang\example\stringutil

```go
package main

import (
    "fmt"
    "github.com/golang/example/stringutil"
)

func main() {
    s := "abc"
    fmt.Println(stringutil.Reverse(s))
}
```

包管理工具

更新依赖
go get -u
go get -u github.com/spf13/hugo
go get -u all

Go代码格式设置
gofmt(用于设置Go文件格式)

golint（提供有关风格和约定方面的建议）
godoc（用于生成和读取文档）

在 Go 语言中，常量与其他数据元素没什么不 同，因此不应将其全大写。

$ gofmt hello.go
$ gofmt -w hello.go

$ go get -u github.com/golang/lint/golint

golint可是找出代码风格方面的问题
$ golint --help
$ go lint hello.go

$ go get golang.org/x/tools/cmd/godoc
$ godoc --help

要给一段代码添加注释，只需在注释行开头指出要注释的元素的名称。

每行都以它注释的类型的名称打头。
注释是以大写字母打头的完整句子，以点结束。
如果对这些代码运行工具 godoc，将生成有关这个包的文档

$ godoc ./hello
$ godoc hello.go

$ godoc strings
$ godoc string | grep "func Replace"

可通过启动一个 Web 服务器来查看标准库文档。在没有连接到网络或连接
速度有限时，这种做法是一个不错的选择。
$ godoc -http=":8080"

 http://localhost:6060/pkg/
 来查看标准库文档

查看包的文档
$ godoc $GOPATH/src/github.com/BurntSushi/toml


单元测试
功能测试
集成测试
基准测试

testing包
```go
package hello

import "testing"

func TestTruth(t *testing.T) {
    if true != true {
        t.Fatal("the world is crumbling")
    }
}
$ go test
```

```go
package hello

import "testing"

type HelloTest struct {
    name string
    locale string
    want string
}

var helloTests = []HelloTest {
    {"a","b","c"},
    {"a","b","c"},
    {"a","b","c"},
}

func TestHello(t *testing.T) {
    for _, test := range helloTests {
        got := Hello(test.name, test.locale)
        if got != test.want {
            t.Errorf("Hello(%s, %s) = %v; want %v", test.name, test.locale, actual, test.want)
        }
    }
}    
```

利用基准测试，测试字符串拼接性能
```go
package hello

import (
    "bytes"
    "strings"
)

func StringFormatAssignment(j int) string {
    var s string
    for i:=0; i<j; i++ {
        s += "a"
    }
    return s
}

func StringformatAppendJoin(j int) string {
    s := []string{}
    for i:=0; i<j; i++ {
        s = append(s, "a")
    }
    return strings.Join(s, "")
    
}

func StringFormatBuffer(j int) string {
    var buffer bytes.Buffer
    for i:=0; i<j; i++ {
        buffer.WriteString("a")
    }
    return buffer.String()
}
```

测试
```go
package hello

import "testing"

func BenchmarkStringFormatAssignment(b *testing.B) {
    for n:=0; n < b.N; n++ {
        StringFormatAssignment(100)
    }
}

func BenchmarkStringFormatAppendJoin(b *testing.B) {
    for n:=0; n < b.N; n++ {
        StringFormatAppendJoin(100)
    }
}

func BenchmarkStringFormatBuffer(b *testing.B) {
    for n:=0; n < b.N; n++ {
        StringFormatAppendJoin(100)
    }
}
```

提供测试覆盖率
$ go test-cover hello.go

使用基准测试
Benchmark
$ go test --bench=.

TDD 测试驱动开发

日志
log包

```go
package main

import (
    "log"
    "errors"
    "os"
)

func main() {
    log.Printf("this is a log message")

    var errFatal = errors.New("something happed")
    log.Fatal(errFatal)

    // 日志写入日志文件
    f, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    log.SetOutput(f)
}
```

日志重定向
$ go run hello.go > log.log 2>&1

打印
```go
package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Guess the name : " )
    text, _ := reader.ReaderString('\n')
    text = strings.Replace(text, "\n", "", -1)

    fmt.Println("[debug] text is ", text)
    ftm.Printf("[debug] text is %v\n", text)
    ftm.Printf("[debug] text is %+v\n", text)
    if test == "John" {
        fmt.Println("you won")
    } else {
        fmt.Println("you are wrong")
    }
}
```

Delve
第三方调试器

$ go get github.com/derekparker/delve/cmd/dlv

$ dlv --help

```go
package main

import "fmt"

func main() {
    s := "abc"
    echo(s)
}
// 
$ dlv debug hello.go
// 设置断点
$ (dlv) break echo
$ (dlv) continue
```

GNU调试器

创建命令行工具

标准输入    0   包含提供给程序的输入
标准输出    1   包含显示到屏幕上的输出
标准错误    2   包含显示到屏幕上的错误

访问命令行参数
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    for i, arg := range os.Args {
        fmt.Println("argument ", i, "is ", arg)
    }
}
```

在 Linux/macOS 系统中，要执行当前工作目录下的可执行文件，必须在指定文件时加上前缀．／

分析命令行标志
```go
package main

import (
    "fmt"
    "flag"
)

func main() {
    s := flag.String("s", "hello world", "string help text")
    i := flag.Int("i", 1, "Int help text")
    b := flag.Bool("b", false, "Bool help text")

    flag.Parse()
    fmt.Println("value of s ", *s)
    fmt.Println("value of i ", *i)
    fmt.Println("value of b ", *b)
}
$ go build hello.go
$ ./hello
$ value of s Hello world

$ ./hello -s Hello baobao
$ value of s Hello baobao

$ ./hello -h
Usage of ./hello:
    -s string
    string help text (default "hello world")

$ ./hello a b c // 跟多个参数

$ ./hello -s Hello -i 11 -b true
```

安装命令行工具

创建web服务器

创建一个基本的web服务器
```go
package main

import (
    "net/http"
)

func hello(w http.ResponseWriter, r *httpRequest) {
    // 设置无响应路径的404
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    // 设置响应报头
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Write([]byte(`{"hello": "world"}`))
}

func main() {
    http.HandleFunc("/", hello) // 创建路由
    http.ListenAndServe(":8080", nil)
}

$ go run hello.go
$ curl -is http://localhost:8080
```

服务端响应类型
text/plain
text/html
application/json
application/xml

根据请求的Accept报头，提供相应的响应类型
```go
package main

import (
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    switch r.Header.Get("Accept") {
        case "application/json":
            w.Header().set("Content-Type", "application/json; charset=utf-8")
            w.Write([]byte(`{"message": "hello"}`))
        case "application/xml":
            w.Header().Set("Content-Type", "application/xml; charset=utf-8")
            w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?><Message>Hello</Message>`))
        default:
            w.Header().Set("Content-Type", "text/plain; charset=utf-8")
            w.Write([]byte("hello"))    
    }
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}

$ curl -si -H 'Accept: application/json' http://localhost:8080
$ curl -si -H 'Accept: application/xml' http://localhost:8080
```

请求端
Http请求类型
GET
POST
PUT
DELETE

响应不同请求类型
```go
package main

import (
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound()
        return
    }

    switch r.Method {
        case "GET":
            w.Write([]byte("receive a GET request\n"))
        case "POST":
            w.Write([]byte("receive a POST request\n"))
        default:
            w.WriteHeader(http.StatusNotImplemented)
            w.Write([]byte(http.StatusText(http.StatusNotImplemented)) + "\n")
    }
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}

$ curl -si -X GET http://localhost:8080
$ curl -si -X POST http://localhost:8080
```

如果请求方法不是 GET 或 POST，就执行 default 子句，发送 501 ( Not Implemented
HTTP ）响应 。 代码 501 意 味着服务器不明白或不支持客户端使用的 HTTP 请求
方法。

https://www.google.com/?q=golang

解析GET请求中的参数类型，以字符串映射的结构来传递
```go
func queryParams(w http.ResponseWriter, r *http.Request) {
    for k, v := range r.URL.Query {
        fmt.Println("%s: %s\n", k, v)
    }
}
```

在POST请求中，数据是在请求体中发送的
```go
func queryParams(w http.ResponseWriter, r *http.Request) {
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", reqBody)
}

```

完整的例子
```go
package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound()
        return
    }

    switch r.Method {
        case "GET":
            for k, v := range r.URL.Query {
                fmt.Printf("%s: %s\n", k, v)
            }
            w.Write([]byte("receive a get request\n"))
        case "POST":
            reqBody, err := ioutil.ReadAll(r.Body)
            if err != nil {
                log.Fatal(err)
            }    
            fmt.Printf("%s\n", reqBody)
            w.Write([]byte("receive a post request\n"))
        default:
            w.WriteHeader(http.StatusNotImplemented)
            w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
    }
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}

$ curl -si "http://localhost:8080/?foo=1&bar=2"
$ curl -si -X POST -d "some data to send" http://localhost:8080/
```

```
$ curl -s -o /dev/null - v http://google.com
> GET I HTT P/1 . 1
> Host : google . com
> Oser- Age nt : curl/7 . 4 3 . 0
> Accept ： 食／＊
>
< HTTP/1 . 1 302 Found
< Cache - Control : private
< Content - Type : text/html; charset=OTF 8
< Referrer - Poli cy : no - referrer
< Location : http : //www . google . co . uk/?gfe_rd=cr&ei=ALMhWdzRK4qwcpmaoLAE

< Con te n t Length : 259
< Date : Su 口 ， 21 May 2017 1 5 : 32 :1 6 GMT
<
{ [2 5 9 byt e s data)
食 Co n nect i on #0 to host google . com l ef t i nt act
```

默认客户端
GET请求
```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    response, err := http.Get("http://ifconfig.io/")
    if err != nil {
        log.Fatal(err)
    }

    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("%s", body)
}

$ go run hello.go
68.235.53.83
```

默认客户端
POST请求
```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

func main() {
    postData := strings.NewReader(`{"some" : "json"}`)
    response, err := http.Post("https://httpbin.org/post", "application/json", postData)
    if err != nil {
        log.Fatal(err)
    }

    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("%s", body)
}
```


自定义客户端
GET请求
```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    client := &http.Client{}
    request, err := http.NewRequest("GET", "http://ifconfig.io/", nil)
    if err != nil {
        log.Fatal(err)
    }

    response, err := client.Do(request)
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("%s", body)
}

$ go run hello.go
68.235.53.83
```

调试请求
```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "os"
    "ioutil"
)

func main() {
    debug := os.Getenv("DEBUG")
    
    client := &http.Client{}
    request, err := http.NewRequest("GET", "https://ifconfig.io", nil)
    if err != nil {
        log.Fatal(err)
    }

    if debug == "1" {
        debugRequest, err := httputil.DumpRequestOut(request, true)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s", debugRequest)
    }

    response, err := client.Do(request)
    defer response.Body.Close()

    if debug == "1" {
        debugResponse, err := httputil.DumpResponse(response, true)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s", debugResponse)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s\n", body)
}
```

在底层，有大量影响响应速度的变数: 
> DNS 查找速度。
> 打开到服务器 IP 地址的 TCP 套接字的速度 。
> 建立 TCP 连接的速度 。
> TLS 握手的速度 （如果连接是 TLS 的） 。
> 向服务器发送数据的速度。
> 重定向的速度 。
> Web 服务器返回响应的速度 。
> 将数据传输到客户端的速度 。

设置请求的超时时间
client := &http.Client {
    Timeout : 1 * time.Second
}

client := &http.Client {
    Timeout : 50 * time.Millisecond
}

通过创建一个传输（ transport ）并将其传递给客户端 ， 可更细致地控制超时 ： 控制 HTTP
连接的各个阶段

tr := &http.Transport{
    DialContext: (&net . Dialer {
    Timeout: 30 * time.Second ,
    KeepAlive: 30 * time.Second,
    }) .D 工 alContext ,
    TLSHandshakeT irneout : 10 * t 工me . Second,
    IdleConnTimeout : 90 * time . Second,
    ResponseHeaderTirneout : 10 * time . Second,
    ExpectContinueTimeout : 1 * time . Second,
}
client := &http.Client {
    Transport : tr,
}


client := &http.Client{}
request, err := http.NewRequest("GET", "http ://www. XXX. com", nil)
request.Header.Add("Connection", "close") // 设置请求的格式

处理json
标准库提供了encoding/json包

JavaScript 对象表示法 (JavaScript Object Notation, JSON ）

结构体转为json

```go
package main

import (
    "fmt"
    "log"
    "encoding/json"
)

type Person struct {
    Name string
    Age int
    Hobbies []string
}

func main() {
    hobbies := []string{"Cycling", "Cheese", "Tech"}
    p := Person {
        Name : "ironman",
        Age : 11,
        Hobbies : hobbies,
    }

    fmt.Printf("%+v\n", p)

    jsonByteData, err := json.Marshall(p)
    if err != nil {
        log.Fatal(err)
    }
    jsonStringData = string(jsonByteData)
    fmt.Println(jsonStringData) 
}
```

使用结构体标签
```go
package main

import (
    "fmt"
    "log"
    "encoding/json"
)

type Person struct {
    Name string `json:"name,omitempty"`
    Age int     `json:"age,omitempty"`
    Hobbies []string    `json:"hobbies,omitempty"`
}

func main() {
    hobbies := []string{"Cycling", "Cheese", "Tech"}
    p := Person {
        Name : "ironman",
        Age : 11,
        Hobbies : hobbies,
    }

    fmt.Printf("%+v\n", p)

    jsonByteData, err := json.Marshall(p)
    if err != nil {
        log.Fatal(err)
    }
    jsonStringData = string(jsonByteData)
    fmt.Println(jsonStringData) 
}
```

json解码
```go
package main

import (
    "fmt"
    "log"
    "encoding/json"
)

type Person struct {
    Name string `json:"name"`
    Age int     `json:"age"`
    Hobbies []string    `json:"hobbies"`
}

func main() {
    jsonStringData := `{"name":"George","age":40,"hobbies": ["Cycling",
"Cheese", "Techno"]}`
    jsonByteData = []byte(jsonStringData)
    p := Person{}
    err := json.Unmarshal(jsonByteData, &p)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", p)
}
```


Go 和 JSON 数据类型
JSON            Go
Boolean         booI
Number          float64
String          string
Array           []interface{}
Object          map[string]interface{}
Null            nil


在json和GO之间映射数据类型
```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Switch struct {
    On bool `json:"on"`
}

func main() {
    jsonStringData := `{"on":"true"}`
    jsonByteData := []byte(jsonStringData)
    s := Switch{}
    err := json.Unmarshal(jsonByteData, &s)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", s)
}
```


通过http获取json

```go
package main

import (
    "net/http"
    "ftm"
    "log"
    "encoding/json"
)

type User struct {
    Name string `json:"name"`
    Blog string `json:"blog"`
}

func main() {
    var u User
    res, err := http.Get("https://api.github.com/users/shapeshed")
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    err = json.NewDecoder(res.Body).Decode(&u)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", u)

}
```

要编码或解码JSON ，必须创建结构体 

如果一个字段可能为空，应给它添加结构体标签 omitempty 。这样解码时，如果该字
段确实为空，将忽略它 。

处理文件
标准库提供了ioutil包

函数 Readfile
将一个文件名作为参数，并以字节切片的方式返回文件的内容

读取文件
```go
package main

import (
    "fmt"
    "log"
    "io/ioutil"
)

func main() {
    fileBytes, err := ioutil.ReadFile("1.txt")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(fileBytes)

    fileString := string(fileBytes)
    fmt.Println(fileString)
}
```

 3 类用户 ： 文件所有者、与文件位于
同一组的用户、其他用户 

函数 WriteFile 接受一个文件名、要写入文件的数据以及应用
于文件的权限

4r  2w  1x

文件权限
符号表示法              数字表示法          说明
----------              0000                无权限
-rwx------              0700                只有所有者能够读取、写入和执行
-rwxrwx---              0770                所有者及其所在的用户组能够读取、写入和执行
-rwxrwxrwx              0777                所有人都能够读取、写入和执行
---x--x--x              0111                所有人都能够执行
--w--w--w-              0222                所有人都能够写入
--wx-wx-wx              0333                所有人都能够写入和执行
-r--r--r--              0444                所有人都能够读取
-r-xr-xr-x              0555                所有人都能够读取和执行
-rw-rw-rw-              0666                所有人都能够读取和写入
-rwxr-----              0740                所有者能够读取、写入和执行，而所有者所在的用户组能够读取

在 UNIX 型系统中 ， 文件的默认权限为 0644，即所有者能够读取和写入，而其他人只能
读取

在文件系统中创建了一个文件 ， 并将其权限设置为 0644
```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    b := make([]byte, 0)
    err := ioutil.WriteFile("1.txt", b, 0644)
    if err != nil {
        log.Fatal(err)
    }

    s := "abc"
    err := ioutil.WriteFile("2.txt", []byte(s), 0644)
    if err != nil {
        log.Fatal(err)
    }

}
```

列出文件目录
类型为 Filelnfo，包含如下信息。
> Name ：文件的名称 。
> Size：文件的长度， 单位为宇节。
> Mode ： 用二进制位表示的权限。
> ModTime：文件最后一个被修改的 时间 。
> IsDir：文件是否是目 录。
> Sys ：底层数据源。


列出目录中的文件及其权限
```go
package main

import (
    "fmt"
    "log"
    "io/ioutil"
)

func main() {
    files ,err := ioutil.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Mode(), file.Name())
    }
}
```

要复制文件，只需结合使用 OS 包中的几个函数。以编程方式复制文件的步骤如下。
1. 打开要复制的文件 。
2. 读取其内容 。
3. 创建井打开要将这些 内 容复制到其 中的文件 。
4. 将内 容写入这个文件 。
5. 关 闭所有己打开的文件 。

将文件的内容复制到一个新文件中

Go 语言致力于确保核心库为小巧而轻量级 的 。 另外，不同的操作性系统差别很大，
这导致创建通用的文件复制方法是很难的。鉴于此，没有用于复制文件的便利方法，要复制
文件，应使用 OS 包 。

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    from, err := os.Open("./1.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer from.Close()

    to, err := os.OpenFile("./1_copy.txt", os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer to.Close()

    _, err = io.Copy(to, from)
    if err != nil {
        log.Fatal(err)
    }
}
```
> 使用 OS 包中的函数 Open 来读取磁盘文件 。
> 使用 defer 语句在程序完成其他所有操作后关闭文件 。
> 使用函数 OpenFile 打开文件 。 第一个参数是要打开（如果不存在 ， 就创建）的文件
的名称 ； 第二个参数是用于文件的标志，在这里指定的是读写文件， 并在文件不存
在时创建它 ； 最后一个参数设置文件的权限 。
> 再次使用 defer 语句在执行完其他操作后关闭文件 。
> 使用 io 包中的函数 Copy 复制源文件的内容，并将其写入 目 标文件 。


删除文件
```go
package main

import (
    "log"
    "os"
)

func main() {
    err := os.Remove("./1.txt")
    if err != nil {
        log.Fatal(err)
    }
}
```

配置文件
```json
{
    "name" : "ironman",
    "awake" : true,
    "hungry" : false
}
```

读取配置文件
```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
)

type Config struct {
    Name string `json:"name"`
    Awake bool `json:"awake"`
    Hungry bool `json:"hungry"`
}

func main() {
    f, err := ioutil.ReadFile("config.json")
    if err != nil {
        log.Fatal(err)
    }

    c := Config{}
    err = json.Unmarshal(f, &c)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", c)
}
```

TOML文件格式
TOML ( Tom’s Obvious, Minimal Language ）是一种专为存储配置文件而设计的格式

解析TOML格式的第三方包
go get github.com/BurntSushi/toml

读取TOML配置文件
```go
package main

import (
    "github.com/BurntSushi/toml"
    "fmt"
    "log"
)

type Config struct {
    Name string
    Awake bool 
    Hungry bool
}

func main() {
    c := Config{}
    _, err := toml.DecodeFile("config.toml", &c)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", c)
}
```


正则表达式
regexp

要在大海里捞针，可使用函数 MatchString
```go
package main

import (
    "fmt"
    "log"
    "regexp"
)

func main() {
    needle := "abc"
    needle1 := "(?i)abc" // 不区分大小写
    haystack := "ssdfsdfsdgsdfsfasf"

    match, err := regexp.MatchString(needle, haystack)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(match) // false
}
```

常用的正则表达式语法
字符                                    含义
.                               与除换行符之前的其他任何字符都匹配
*                               与零个或多个指定的字符匹配
^                               表示行首
$                               表示行尾
+                               匹配一次或多次
?                               匹配零或一次
[]                              与方摇号内指定的任何字符都匹配
{n}                             匹配 n 次
{n,}                            匹配 n 次或更多次
{m,n}                           最少匹配 m 次，最多匹配 n 次

eg:

要检查的条件
> 应跃于 4 个字符 ， 但不超过 12 个字符 。
> 应只包含字母和数字 。
> 字符可大写 ，也可小写。

^[a-zA-Z0-9]{5,12}$

> 字符^表示从字符串开头开始匹配。
> 方括号（ ［］ 〉内的字符集表示与其中的任何字符都匹配。
> 大括号（｛）〉 内的数字表示应至少匹配 5 次 ，但最多不超过 12 次。


用于分析正则表达式的函数有两个
> Compile：在正则表达式未能通过编译时返回错误
> MustCompile ：在正则表达式无法编译时引发 panic
该使用哪一个取决于具体情况，但 MustCompile 通常是更佳的选择。

当正则表达式无法分析时，函数 Compile 返回错误，而函数 Mus tCompile 引发 panic 
标志 i 让正则表达式不区 分大小 写

reg := regexp.MustCompile("^[a-zA-Z0-9]{5,12}$")


```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    reg := regexp.MustCompile("^[a-zA-Z0-9]{5,12}$")
    fmt.Println(regexp.MatchString("slimshady99") // true
}
```

reg1 := regexp.MustCompile("^[a-zA-Z0-9]{5,12}$")
reg2 := regexp.MustCompile("^[a-zA-Z0-9]{5,12}$")
if !reg1.MatchString(a) {
    a = reg2.ReplaceAllString(a, "x")
}


时间
time包

打印当前系统时间
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println(time.Now())
}
```

Linux设置系统时间 
sudo date +%Y%m%d -s "20500101"

网络时间协议（ Network Time Protocol,
NTP ）是一种在整个网络中同步时间的网络协议， 使用 NTP 的不同计算机更有可能就时间达
成一致，但在本地它们依然可以设置不同的时区


在计算中 ，要消除时 区的影响 ，可参照世界标准时间（ Coordinated Universal Time, UTC ） 。
UTC 是时间标准而非时区，它让不同地区的计算机有相同的参照物 ，而不用考虑相对时区

时间休眠
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    time.Sleep(3 * time.Second)
    fmt.Println("i'm awake")
}
```

设置超时
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("you have two seconds to caculate ... ")

    for {
        select {
            case <-time.After(2 * time.Second):
                fmt.Println("Time' up! The answer is 1. Did you get it?")
                return
        }
    }
}
```

定时时间
ticker

```go
package main

import (
    "ftm"
    "time"
)

func main() {
    c := time.Tick(5 * time.Second)
    for t := range c {
        fmt.Printf("the time is now %v\n", t)
    }

}
```


时间的字符串表示
类型        字符串
ANSIC       Mon Jan 2 15:04:05 2006
UnixDate    Mon Jan 2 15:04:05 MST 2006
RubyDate    Mon Jan 02 15:04:05 -0700 2006
RFC822      02 Jan 06 15:04 MST
RFC822Z     02 Jan 06 15:04 -0700
RFC850      Monday, 02-Jan-06 15:04:05 MST
RFCll23     Mon, 02 Jan 2006 15:04:05 MST
RFCll23Z    Mon, 02 Jan 2006 15:04:05 -0700
RFC3339     2006-0 l-02T I5:04:05Z07:00
RFC3339Nano 2006-0 I-02T I5:04:05.999999999Z07:00

```go
package main

import (
    "ftm"
    "time"
)

func main() {
    s := "2019-01-01T11:11:11+07:00"
    t, err := time.Parse(time.RFC3339, s)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(err)

}
```

指定平台进行编译

要为 32 位的 Linux系统编译程序
GOOS=linux GOPATH=386 go build hello.go

要为 64 位的 Windows 系统编译程序
GOOS=windows GOPATH=amd64 go build hello.go


Go 编程语言的优点之一是 ， 支持大量的操作系统和体系结构组合 ，因此程序员可为这些
不同的目标编译 Go 代码。

压缩生成的二进制文件大小
通过指定一些编译标志，可压缩编译得到的二进制文件的大小 。 这些标志指定省略符号
表、调试信息和 ：DWARF 符号表。

GOOS=linux go build -ldflags="-s -w" hello.go

Go 官网维护着一些 Docker 映像，让您能够编译代码并在 Do cker 容器中运行它们 。


常量

Go 语言预定义了这些常量：true、false 和 iota。

iota 比较特殊，可以被认为是一个可被编译器修改的常量，在每一个 const 关键字出现时被重置为 0，然后在下一个 const 出现之前，每出现一次 iota，其所代表的数字会自动增 1。


```go
package main

const (    // iota 被重置为 0
    c0 = iota   // c0 = 0
    c1 = iota   // c1 = 1
    c2 = iota   // c2 = 2
)

const (
    u = iota * 2;  // u = 0
    v = iota * 2;  // v = 2
    w = iota * 2;  // w = 4
)

const x = iota;  // x = 0
const y = iota;  // y = 0
```

如果两个 const 的赋值语句的表达式是一样的，那么还可以省略后一个赋值表达式。因此，上面的前两个 const 语句可简写为：

```go
const ( 
    c0 = iota 
    c1 
    c2 
)

const ( 
    u = iota * 2 
    v 
    w 
)
```

常量还可以用于枚举。
枚举中包含了一系列相关的常量，比如下面关于一个星期中每天的定义。Go 语言并不支持其他语言用于表示枚举的 enum 关键字，而是通过在 const 后跟一对圆括号定义一组常量的方式来实现枚举。

下面是一个常规的 Go 语言枚举表示法，其中定义了一系列整型常量：
```go
const (
    Sunday = iota 
    Monday 
    Tuesday 
    Wednesday 
    Thursday 
    Friday 
    Saturday 
    numberOfDays
)
```

基本数据类型：
布尔类型：bool
整型：int8、byte、int16、int、uint、uintptr 等
浮点类型：float32、float64
复数类型：complex64、complex128
字符串：string
字符类型：rune
错误类型：error


复合类型：
指针（pointer）
数组（array）
切片（slice）
字典（map）
通道（chan）
结构体（struct）
接口（interface）

```go
var v1 bool 
v1 = true 
v2 := (1 == 2) // v2 也会被推导为 bool 类型

var b bool 
b = 1 // 编译错误 
b = bool(1) // 编译错误

var b bool 
b = (1!=0) // 编译正确 
fmt.Println("Result:", b) // 打印结果为Result: true

```

! 运算符也不能作用于非布尔类型值

不同类型的值直接不能使用 == 或 != 运算符进行比较，在编译期就会报错
b := (false == 0);

int 和 int32 在 Go 语言里被认为是两种不同的类型（同理，int 和 int64 也是不同的类型），编译器也不会帮你自动做类型转换，比如以下的例子会有编译错误：
var int_value_1 int8
int_value_2 := 8   // int_value_2 将会被自动推导为 int 类型 
int_value_1 = int_value_2  // 编译错误
int_value_1 = int8(int_value_2)) // 编译通过

可以通过增加前缀 0 来表示八进制数（如：077），增加前缀 0x 来表示十六进制数（如：0xFF），以及使用 E 来表示 10 的连乘（如：1E3 = 1000

不同类型的整型值不能直接进行算术运算

不同类型的值不能放在一起比较，否则会报编译错处，不过，各种类型的整型变量都可以直接与字面常量进行比较，比如：
if int_value_1 == 8 {
    fmt.Println("int_value_1 = 8")
}

运算符优先级如下所示（由上到下表示优先级从高到低，或者数字越大，优先级越高）：

6      ^（按位取反） !
5      *  /  %  <<  >>  &  &^
4      +  -  |  ^（按位异或）
3      ==  !=  <  <=  >  >=
2      &&
1      ||

注：++ 或 -- 只能出现在语句中，不能用于表达式，故不参与优先级判断

在 Go 语言里，定义一个浮点数变量的代码如下：

var float_value_1 float32

float_value_1 = 10
float_value_2 := 10.0 // 如果不加小数点，float_value_2 会被推导为整型而不是浮点型
float_value_3 := 1.1E-10

对于浮点类型需要被自动推导的变量，其类型将被自动设置为 float64，而不管赋值给它的数字是否是用 32 位长度表示的。

float_value_1 = float32(float_value_2)

在实际开发中，应该尽可能地使用 float64 类型，因为 math 包中所有有关数学运算的函数都会要求接收这个类型。

浮点数的运算和整型一样，也要保证操作数的类型一致，float32 和 float64 类型数据不能混合运算

p := 0.00001
// 判断 float_vlalue_1 与 float_value_2 是否相等
if math.Dim(float64(float_value_1), float_value_2) < p {
    fmt.Println("float_value_1 和 float_value_2 相等")
} 


bin 编译后的可执行文件
pkg 编译后的包文件
src 源码


$ go doc goWorkspace
$ go doc ftm
$ godoc fmt Println
$ godoc -http=:8080

$ go run 直接运行程序
$ go build 测试编译，检查是否有编译错误
$ go fmt 格式化源码
$ go install 编译包文件并编译整个程序
$ go test 运行测试文件








