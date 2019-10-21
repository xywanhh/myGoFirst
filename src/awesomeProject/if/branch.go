package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %s", score))
	case score < 70:
		g = "F"
	case score > 80:
		g = "A"
	default:
		g = "B"
	}
	return g
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileName string) {
	// 一行行的读取文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

const fileName = "E:\\githubRepositorys\\myGoFirst\\src\\awesomeProject\\if\\1.txt"

func main() {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	// 进阶写法
	if contents, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(grade(100))
	fmt.Println(grade(50))
	fmt.Println(grade(90))
	//grade(110)

	fmt.Println(convertToBin(5))

	printFile(fileName)
}
