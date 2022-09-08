package main

import "fmt"

func Func1(a int, b string, c bool) {
	fmt.Println("testing", a, b, c)
}

func Func2() (int, string, bool) {
	return 0, "a", true
}

func Func3() (string, bool) {
	return "b", false
}

func main() {
	Func1(123, "abc", true)
	Func1(Func2())
	// compile error
	// Func1(123, Func3())
}
