package main

import "fmt"

type P struct {
	x, y int
}

type TreeNode struct {
	Left, Right *TreeNode
	Value int
}

// 为结构体定义方法
func (node *TreeNode) print() {
	fmt.Println(node.Value)
}

func (node *TreeNode) setValue(v int) {
	if node == nil {
		fmt.Println("setting value to nil" + "node Ignored")
		return
	}
	node.Value = v
}

// 遍历
func (node *TreeNode) traverse() {
	if (node == nil) {
		return
	}
	node.Left.traverse()
	node.print()
	node.Right.traverse()
}

// 工厂函数
func createNode(value int) *TreeNode {
	return &TreeNode{Value : value}
}

// 扩展TreeNode的功能
type MyTreeNode struct {
	node *TreeNode
}

func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	leftNode := MyTreeNode{myNode.node.Left}
	leftNode.postOrder()
	rightNode := MyTreeNode{myNode.node.Right}
	rightNode.postOrder()
	myNode.node.print()
}

func main() {
	var root1 TreeNode
	fmt.Println(root1)

	root2 := TreeNode{Value : 3}
	root2.Left = &TreeNode{}
	root2.Right = &TreeNode{nil, nil, 5}
	root2.Right.Left = new(TreeNode)
	root2.Left.Right = createNode(2)
	root2.Right.Left.setValue(4)
	fmt.Println(root2)

	fmt.Println()
	root2.traverse()

	fmt.Println()
	myRoot := MyTreeNode{&root2}
	myRoot.postOrder()

}
