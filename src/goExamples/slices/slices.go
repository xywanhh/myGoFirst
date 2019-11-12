package main

import "fmt"

/*
Unlike arrays, slices are typed only by the elements they contain (not the number of elements).

*/
func main()  {
	// To create an empty slice with non-zero length, use the builtin make.
	s := make([]string, 3)
	fmt.Println("empty:", s)
	
	// set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get: ", s[0])
	fmt.Println("set: ", s)

	// len returns the length of the slice as expected.
	fmt.Println(len(s))

	// builtin append, which returns a slice containing one or more new values. 
	// Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append: ", s)

	// Slices can also be copy’d.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	// slice[low:high]. 

	// elements s[2], s[3], and s[4].
	l1 := s[2:5]
	fmt.Println("l1: ", l1)

	// This slices up to (but excluding) s[5]. 
	l2 := s[:5]
	fmt.Println("l2: ", l2)

	l3 := s[2:]
	fmt.Println("l3: ", l3)
	
	// declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "l"}
	fmt.Println("declare: ", t)

	// Slices can be composed into multi-dimensional data structures. 
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i:=0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j:=0;j<innerLen;j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)


}