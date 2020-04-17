package main

import (
	"fmt"
	"regexp"
	//"bufio"
)

func main() {
	//1. string is a byte array
	str := "This is ä string."
	fmt.Println("byte length: ", len(str))

	strRunes := []rune(str)
	fmt.Println("rune length: ", len(strRunes))

	firstRune, lastRune := strRunes[0:1], strRunes[len(strRunes)-1:]

	fmt.Println("first character:", string(firstRune))
	fmt.Println("last character:", string(lastRune))

	//2. Delete symbols with count > 1
	m := make(map[rune]int)

	for _, k := range str {
		if v, ok := m[k]; ok {
			m[k] = v + 1
		} else {
			m[k] = 1
		}
	}

	for _, k := range str {
		if m[k] > 1 {
			continue
		}
		fmt.Print(string(k))
	}
	fmt.Println()

	//3. Password check
	password := "haius85QQWEs"
	isStrong := CheckPassword(password)

	fmt.Println(isStrong)
}

func CheckPassword(str string) bool {
	rx := regexp.MustCompile("^[a-zA-Z0-9]*$")

	if len(str) < 5 {
		return false
	}
	if !rx.MatchString(str) {
		return false
	}

	return true
}
