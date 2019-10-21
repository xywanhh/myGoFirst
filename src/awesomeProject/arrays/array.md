## 数组
```text
数组是值类型
[10]int 和 [5]int是不同的类型
调用func f(arr [10]int)会拷贝数组
一般不使用数组
```

## 切片slice
```text
slice本身没有数据，是对底层数组的一个view

slice的实现
ptr --> 指向数组的头一个元素
array --> 底层数组
len --> 长度
cap --> 容量

slice可以向后扩展，不可以向前扩展

s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)

s = append(s, 1)
s1 := append(s, 1) // 扩展

向slice添加元素
- 添加元素时如果超过cap，系统会重新分配更大的底层数组
- 由于值传递的关系，必须接收append的返回值

通过make创建slice
s1 := make([]int, 10)
s2 := make([]int, 10, 32)

slice的拷贝
copy(s2, s1)

没有内建的删除
s2 = append(s2[:3], s2[4:]...)
```
