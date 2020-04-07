package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func main() {
	var foo int

	foo = 42

	var bar int = 'c'

	fmt.Println(foo)
	fmt.Println(string(bar))
	fmt.Println(string(foo))

	//like C#!

	var a = "Hello, Jo!"

	fmt.Println(a)

	//short hand

	short := "I'm long!"

	fmt.Println(short)

	//block
	var (
		x1 = 123
		x2 = 321
	)

	fmt.Println(x1, x2)

	//read from std.input

	//not working in VS Code
	//fmt.Scan(&a)
	//fmt.Println(a)

	z := 12
	var asd string

	asd = strconv.Itoa(z) + "asd"

	fmt.Println(asd[len(asd)-2 : len(asd)-1])

	var d int = 105

	hours := d / 30
	foo2 := d % 30
	minutes := foo2 * 2

	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)

	// 1
	var k int = 2
	//fmt.Scan(&k)
	if k == 0 {
		fmt.Println("Ноль")
	} else if k%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}

	//2
	var l int = 2
	//fmt.Scan(&k)
	if l == 0 {
		fmt.Println("Ноль")
	} else if l > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число отрицательное")
	}

	//3
	var m int = 1342

	strm := strconv.Itoa(m)
	fmt.Println(strm[1])

	var xx string
	xx = strm[1:]
	fmt.Println(xx)

	const nihongo = "日本語gg"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	//4
	fmt.Println(isUnique(1232))

	//5
	fmt.Printf("%q\n", strm[0:1])

	foo2, _ = strconv.Atoi(strm[0:1])
	fmt.Println(foo2)

	//6
	fmt.Println(isHappy(123321))

	//7
	if m == 0 {
		return
	} else if m > 0 {
		fmt.Printf("Positive. %d %d \n", m, m+1)
		fmt.Printf("Positive. %d %d \n", m, m-1)
	} else if m < 0 {
		fmt.Printf("Negative. %d %d \n", m, m+1)
		fmt.Printf("Negative. %d %d \n", m, m-1)
	}

	//8
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}

	//9
	bank()

	//10
	fmt.Println("---------------- 10 --------------")
	var aaa, bbb int = 5432, 4311
	var res string = ""
	for _, val := range strconv.Itoa(aaa) {
		for _, val2 := range strconv.Itoa(bbb) {
			if val == val2 {
				res += string(val)
			}
		}
	}

	for i := 0; i < len(res); i++ {
		fmt.Print(res[i:i+1], " ")
	}

	//11
	fmt.Println("\n---------------- 11 --------------")

	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	//удаляем 4-й элемент
	var n = 3
	users = append(users[:n], users[n+1:]...)
	fmt.Println(users) //["Bob", "Alice", "Kate", "Tom", "Paul", "Mike", "Robert"]

	//12
	v := '1'
	fmt.Println(strconv.ParseInt(string(v), 10, 10))
	fmt.Println(strconv.Atoi(string(v)))

	fmt.Println(reflect.TypeOf(v))
}

func isUnique(number int) bool {
	str := strconv.Itoa(number)

	for i := 0; i < len(str); i++ {
		count := 0
		for j := 0; j < len(str); j++ {
			if str[j] == str[i] {
				count++
			}
		}
		if count > 1 {
			return false
		}
	}

	return true
}

func isHappy(number int) bool {
	str := strconv.Itoa(number)

	var sum1 rune = 0
	for _, val := range str[0:3] {
		sum1 += val
	}
	//fmt.Println(sum1)

	var sum2 rune = 0
	for _, val := range str[3:6] {
		sum2 += val
	}
	//fmt.Println(sum2)

	return sum1 == sum2
}

func isLeap(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}

	return false
}

func bank() {
	var x, p, y float64 = 100, 10, 200

	count := 1
	var percent float64
	var sum float64 = x

	for 1 == 1 {
		percent = sum * p / 100
		sum += percent
		sum = math.Round(sum)

		if sum >= y {
			break
		}

		count++
	}

	fmt.Println(count)
}
