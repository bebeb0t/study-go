package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func execFunc() {
	fmt.Println(add(42, 13))
}
