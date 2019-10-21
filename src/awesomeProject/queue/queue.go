package main

import "fmt"

// 利用组合扩展slice的功能

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func main() {
	q := Queue{1, 2}
	q.Push(3)
	q.Push(4)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.isEmpty())
}
