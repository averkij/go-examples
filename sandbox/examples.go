package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

//small tasks from the Stepik Go course and other sources

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
	fmt.Println("\n---------------- 12 --------------")
	v := '1'
	fmt.Println(strconv.ParseInt(string(v), 10, 10))
	fmt.Println(strconv.Atoi(string(v)))

	fmt.Println(reflect.TypeOf(v))

	//13. reverse number
	fmt.Println("\n---------------- 13 --------------")
	var a22 string = "3755"

	for i := len(a22) - 1; i >= 0; i-- {
		fmt.Print(string(a22[i]))
	}

	//14. int to hours and secons
	fmt.Println("\n---------------- 14 --------------")
	var ksec int = 13257

	fmt.Printf("It is %d hours %d minutes.\n", ksec/3600, ksec%3600/60)

	//15
	fmt.Println("\n---------------- 15 --------------")
	var a15, b15, c15 float64 = 6, 8, 10

	if math.Sqrt(a15*a15+b15*b15) == c15 {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}

	fmt.Println("\n---------------- 16 --------------")
	strToSum := 123123222

	for strToSum > 9 {
		strToSum = getSum(strToSum)
	}

	fmt.Println(getSum(strToSum))

	//17
	fmt.Println("\n---------------- 17 --------------")
	for i := 0; i < 100; i++ {
		if i%10 == 1 && i != 11 {
			fmt.Println(i, "korova")
		} else if (i%10 == 2 || i%10 == 3 || i%10 == 4) && i != 12 && i != 13 && i != 14 {
			fmt.Println(i, "korovy")
		} else {
			fmt.Println(i, "korov")
		}
	}

	//18
	fmt.Println("\n---------------- 18 --------------")
	a18 := 256
	k18 := 1
	for k18 <= a18 {
		fmt.Print(k18, " ")
		k18 = k18 << 1
	}

	fmt.Println("\n---------------- 19 --------------")
	fiboLimit := 8

	for i := range getFibos(fiboLimit) {
		fmt.Println(i)
	}

	//20
	num20 := 123
	str20 := strconv.Itoa(num20)
	fmt.Println("number of digits:", len(str20))

	//21 pointer
	v21 := 5
	p21 := &v21
	fmt.Print(*p21, " ")
	changePointer(p21)
	fmt.Print(*p21)
}

func changePointer(p *int) {
	v := 3
	p = &v
}

func getFibos(limit int) <-chan int {
	ch := make(chan int)

	go func() {
		a, b := 0, 1
		for a <= limit {
			ch <- a
			b = a + b
			a = b - a
		}
		close(ch)
	}()

	return ch
}

func getSum(num int) int {
	str := strconv.Itoa(num)
	sum := 0
	for _, v := range str {
		val, _ := strconv.Atoi(string(v))
		sum += val
	}
	return sum
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
