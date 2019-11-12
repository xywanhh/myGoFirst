package main

import "fmt"

/*
 Assigning a value to a dereferenced pointer changes the value at the referenced address.
*/

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func main()  {
	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroval:", i)

	zeroPtr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

}