package main

import "fmt"

func execTypeInference() {
	v := 0.867 + 0.5i
	fmt.Printf("v is of type %T %v\n", v, v)
}
