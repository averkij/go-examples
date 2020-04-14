package main

import (
	"fmt"
	"reflect"
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

	//3
	typeRefl := reflect.TypeOf(obj)
	valRefl := reflect.ValueOf(obj)
	fmt.Println(typeRefl, valRefl)
}
