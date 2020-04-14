package main

import (
	"fmt"
)

func main() {
	var obj interface{} = []int {1,2}

	//1
	switch t := obj.(type) {
	case int:
		fmt.Println("int:", t)
	case []int:
		fmt.Println("int[]:", t)
	default:
		fmt.Println("unknown")
	}

	//2
	typeName := fmt.Sprintf("%T", obj)
	fmt.Println(typeName)
}
