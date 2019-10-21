## map
```text
创建：
- make(map[string]string)
- var map[string]string
- m := map[string]string{}

获取元素：
m[key] // key可以不存在
key不存在的时候，获取value类型的初始值

用value, ok := m[key]来判断key是否存在

用delete删除一个key

map遍历
range
不保证遍历顺序
使用len获取元素个数

map的key
map使用哈希表，必须可以比较相等
除了slice,map,function的内建类型都可以作为key
struct类型不包含上述字段，也可以作为key
```
```text
使用range遍历pos,rune对
使用utf8.RuneCountInString获得字符数量
使用len获得字节长度
使用[]byte获得字节
```