## 传统测试和表格驱动测试
```text
// 传统
@Test
public void testAdd() {
    assertEquals(3, add(1, 2));
    assertEquals(3, add(0, 2));
    assertEquals(3, add(0, 0));
    assertEquals(3, add(-1, 2));
    assertEquals(Integer.MIN_VALUE, add(1, Integer.MAX_VALUE));
// 传统缺点
- 测试数据和测试逻辑混合在一起
- 出错信息不明显
- 一旦一个数据出错，所有测试全部结束

// go的测试理念：表格驱动测试
**TestAdd()**

tests := []struct {
    a, b, c int32
} {
    {1, 2, 3},
    {1, 2, 3},
    {1, 2, 3},
    {math.MaxInt32, 1, math.MinInt32},
}
for _, tt := range tests {
    if actual := add(tt.a, tt.b); actual != tt.c {
        t.Errorf("calc(%d, %d); " + 
            "got %d; expected %d", 
            tt.a, tt.b, actual, tt.c)
    }
}
// 表格驱动的有点：
- 分离的测试数据和测试逻辑
- 明确的出错信息
- 额可以部分失败


// 代码覆盖率和性能测试
$ go test .
$ go test fmt

**BenchmarkAdd()**
$ go test -bench

生成性能报告文
$ go test -bench . -cpuprofile cpu.out
$ go tool pprof cpu.out
$ web

b *testing.B
b.Logf(len(s))
b.ResetTimer()

// 总结
testing.T的使用
运行测试
使用IDEc查看你diamagnetic覆盖
使用go test获取代码覆盖报告
使用go tool cover 查看代码覆盖报告

性能测试
testing.B的使用
使用pprof优化性能


http测试
- 通过假的Request/Response
- 通过起夫区其
}
```

## 生成文档和示例代码
```text
$ go doc 
$ go doc fmt
$ go doc fmt.Println
$ godoc -http :6060

// 注释语句
// e.g: fmt.Println("hello")

生成示例代码
import "fmt"
func ExampleAdd() {
    fmt.Println(1)
}

// Output:
// 1
// 2
// false

文档
- 用注释协文档
- 在测试中加入Example
- 使用go doc/godoc来查看，生成文档











```

