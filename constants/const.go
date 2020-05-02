package main

import (
	"fmt"
)

func main() {
	const (
		a = iota + 1
		_
		b
		c
		d = c + 2
		t
		i
		i2 = iota + 2
	)

	fmt.Println(a) //1
	fmt.Println(b) //3
	fmt.Println(c) //4
	fmt.Println(d) //6
	fmt.Println(t) //6
	fmt.Println(i) //6!
	fmt.Println(i2) //9
}