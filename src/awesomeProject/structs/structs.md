## 结构体
```text
创建
root := TreeNode{Value : 1}
root.Left = &TreeNode{}
root.Right = &TreeNode{nil, nil, 100}
root.Right.Left = new(TreeNode)

不论是结构体还是地址，一律使用.来访问成员

// 为结构体定义方法
func (node *TreeNode) print() {
	fmt.Println(node.Value)
}
nil指针也可以调用方法
```

## 封装
```text
名字一般CamelCase
首字母大写：public
首字母小写：private

为结构体定义的函数必须放在一个包内
一个包内只能有一个main函数入口
```

## 扩展系统类型或者别人的包
- 定义别名
- 使用组合

```text
go install 产生pkg文件和可执行文件
go run 直接编译运行
go build 来编译
```