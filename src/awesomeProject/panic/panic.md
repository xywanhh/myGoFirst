## panic
```text
停止当前函数执行
一直向上执行，执行每一层的defer
如果没有遇到recover,程序退出

```

## recover
```text
仅在defer调用中使用
获取panic的值
如果无法处理，可以重新panic
```