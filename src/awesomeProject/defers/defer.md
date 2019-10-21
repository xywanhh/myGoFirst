## defer

```text
确保在函数结束时发生

参数在defer语句时计算

defer列表为先进后出

```

## error
```text
file, err := os.Openfile(filename, os.O_EXCL|os.O_CREATE, 0666)
err := errors.New("something happened") // 自定义error
if err != nil {
    if pathError, ok := err.(*os.PathError); !ok {
        panic(err)
    } else {
        fmt.Printf("%s, %s, %s\n", 
            pathError.Op,  
            pathError.Path,
            pathError.Err)
    }
}
```

```go
type error interface {
    Error() string
}
```