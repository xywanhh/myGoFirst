package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile(filename string)  {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := func() string {
		return "abc"
	}
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("ff.txt")
}
