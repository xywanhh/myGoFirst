## 内建类型：
- bool
- string
- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
- byte, rune
- float32, float64, complex64, complex128

## 强制类型转换
> a = int(b)

## 常量
> const 关键字<br/>
> 常量值可以当作任意类型的值使用
```go
const a, b = 3, 4 // 没有定义类型
int c = int(math.Sqrt(a*a+b*b))
```

## go语言没有枚举类型
> 使用常量来实现枚举
- 普通枚举类型 （const）
- 自增枚举类型 （iota）