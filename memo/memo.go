package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Memory trainer
//
//1. At first build the numbers to persons table according to the "Quantum memory power" book.
//2. Train this table.
//3. Put the persons to your memory palace.
//4. Recall the numbers from the palace.

func main() {
	intervalInSeconds := 4
	amountOfNumbers := 20

	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)

	for i := 0; i < amountOfNumbers; i++ {
		n := rnd.Int31n(100)
		if n < 10 {
			fmt.Print("0")
		}
		fmt.Print(n, " ")
		time.Sleep(time.Second)

		if i != amountOfNumbers-1 {
			if (i+1)%10 != 0 {
				for j := 0; j < intervalInSeconds-1; j++ {
					fmt.Print("-")
					time.Sleep(time.Second)
				}
				fmt.Print(" ")
			} else {
				fmt.Print("\n\n")
			}
		}
	}
}
