package main

import (
	"fmt"
	"time"
)

func Wrong(m map[int]int) {
	fmt.Println("wrong:")
	for k, v := range m {
		go func() {
			// loop variable k captured by func literal
			fmt.Println("k: ", k, "v: ", v)
		}()
	}
}

func Correct(m map[int]int) {
	fmt.Println("correct:")
	for k, v := range m {
		go func(a, b int) {
			fmt.Println("a: ", a, "b: ", b)
		}(k, v)
	}
}

func main() {
	m := make(map[int]int, 10)
	for i := 1; i <= 10; i++ {
		m[i] = i
	}

	Wrong(m)
	time.Sleep(time.Second)

	Correct(m)
	time.Sleep(time.Second)
}
