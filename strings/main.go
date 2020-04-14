package main

import (
	"fmt"
	//"bufio"
)

func main() {
	str := "This is ä string."
	fmt.Println("byte length: ", len(str))

	strRunes := []rune(str)
	fmt.Println("rune length: ", len(strRunes))

	firstRune, lastRune := strRunes[0:1], strRunes[len(strRunes)-1:]

	fmt.Println("first character:", string(firstRune))
	fmt.Println("last character:", string(lastRune))
}
