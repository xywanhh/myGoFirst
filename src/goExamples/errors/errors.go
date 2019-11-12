package main

import "fmt"
import "errors"

/*
errors are the last return value and have type error, a built-in interface.

errors.New constructs a basic error value with the given error message.
errors.New 构造一个使用给定的错误信息的基本error值。

返回错误值为 nil 代表没有错误。

通过实现 Error 方法来自定义 error 类型

注意在 if行内的错误检查代码，在 Go 中是一个普遍的用法。
*/

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg int
	prob string
}

// 实现内置的error接口 Error() string
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main()  {

	for _, i := range []int{1, 2, 42} {
		if r, err := f1(i); err != nil {
			fmt.Println("f1 fail: ", err)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}

	for _, i := range []int{1, 2, 42} {
		if r, err := f2(i); err != nil {
			fmt.Println("f2 fail: ", err)
		} else {
			fmt.Println("f2 worked: ", r)
		}
	}

	// 如果想在程序中使用一个自定义错误类型中的数据，需要通过类型断言来得到这个错误类型的实例。
	_, err := f2(42)
	if t, ok := err.(*argError); ok {
		fmt.Println(t.arg)
		fmt.Println(t.prob)
	}
}