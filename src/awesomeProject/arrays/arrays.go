package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("len:%d cap:%d\n", len(s), cap(s))
}

func main() {
	// 数组声明
	var arr1 [2]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 3, 4}
	var arr4 [4][2]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(arr4)

	// 数组遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// 数组遍历 range
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}

	// 求最大值
	maxi := -1
	maxValue := -1
	for i, v := range arr3 {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	fmt.Println(maxi, maxValue)

	// sum
	sum := 0
	for _, v := range arr3 {
		sum += v
	}
	fmt.Println(sum)

	arr := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(arr, s1, s2)

	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i * 2 + 1)
	}
	fmt.Println(s)

	fmt.Println("copying slice")
	copy(s2, s1)
	fmt.Println(s2)

	fmt.Println("deleting elements")
	s1 = append(s1[:2], s1[3:]...)
	fmt.Println(s1)


}
