package main

import (
	"awesomeProject/duck"
	"fmt"
)

type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get("http://www.baidu.com")
}

type Poster interface {
	Post(url string, form map[string]string)
}

func post(p Poster)  {
	p.Post("http://www.baidu.com",
		map[string]string{
			"name" : "abc",
		})
}

type XinP interface {
	Retriver
	Poster
}

const url  = "http://www.baidu.com"
func session(p XinP) string {
	p.Post(url, map[string]string{
		"name" : "123",
	})
	return p.Get(url)
}

func main() {
	var r Retriver

	//fmt.Println(download(r))

	//r = duck.Xin{"this is fake"}
	//fmt.Println(download(r))

	r = duck.Xin1{}
	fmt.Println("%T %v\n", r, r)
	//fmt.Println(download(r))
	
	// 查看接口里的类型
	switch  v := r.(type) {
	case duck.Xin1:
		fmt.Println("contents: ", v.TimeOut)
	case duck.Xin:
		fmt.Println("TimeOut: ", v.Contents)
	}

	// Type assertion
	r1 := r.(duck.Xin1) // 取得接口里的真正类型
	fmt.Println(r1.TimeOut)
}
