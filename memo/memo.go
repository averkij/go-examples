package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	intervalInSeconds := 2

	for i := 0; i < 10; i++ {
		fmt.Print(rnd.Int31n(100))
		time.Sleep(time.Second * time.Duration(intervalInSeconds))
		fmt.Print(" - ")
		time.Sleep(time.Second)
	}
}
