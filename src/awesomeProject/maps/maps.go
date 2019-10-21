package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	maps := map[string]string {
		"name" : "abc",
		"age" : "11",
	}
	fmt.Println(maps)
	for k, v := range maps {
		fmt.Println(k, v)
	}
	fmt.Println(maps["name"])
	fmt.Println(maps["name1"])

	name, ok := maps["name"]
	fmt.Println(name, ok)
	name1, ok := maps["name1"]
	fmt.Println(name1, ok)

	if name, ok := maps["name1"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("key not exists")
	}

	m2 := make(map[string]int)
	fmt.Println(m2) // m2 === empty map

	var m3 map[string]string
	fmt.Println(m3) // m3 === nil

	// 寻找最长不包含重复字符的子串
	// abcabcaa -> abc
	fmt.Println(lengthOfNonRepeatingSubstr("abcabcaaabcdef"))
	fmt.Println(lengthOfNonRepeatingSubstr("abcabcaa"))

	fmt.Println("aa->", lengthOfNonRepeatingSubstr1("yes我爱"))
	fmt.Println("bb->", lengthOfNonRepeatingSubstr("yes我爱"))

	s10 := "Yes我爱北京台！" // utf-8编码格式，中文3字节
	fmt.Println(len(s10))
	fmt.Printf("%X\n", []byte(s10))
	for _, b := range []byte(s10) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s10 { // ch is rune
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println(utf8.RuneCountInString(s10))

	bytes := []byte(s10)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s10) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}

func lengthOfNonRepeatingSubstr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func lengthOfNonRepeatingSubstr1(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

// rune 相当于go里的char
