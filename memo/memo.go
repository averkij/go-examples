package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

//Memory trainer
//
//1. Build the numbers to persons table according to the "Quantum memory power" book.
//2. Train this table.
//3. Put the persons to your memory palace.
//4. Recall the persons from the palace and translate them to numbers.

func main() {
	//cards training
	trainCards(52, 2)

	//memory palace training
	reverse := false
	random := false
	showTag := false
	showNumber := true
	oneLine := true
	trainPalace(2, reverse, random, showTag, showNumber, oneLine)
}

func trainCards(cardsAmount, intervalInSeconds int) {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	fmt.Print("\n")

	for i := 0; i < cardsAmount; i++ {
		n := rnd.Int31n(100)
		if n < 10 {
			fmt.Print("0")
		}
		fmt.Print(n, " ")
		time.Sleep(time.Second)

		if i != cardsAmount-1 {
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

//52 cells for cards memorising
func trainPalace(intervalInSeconds int, reverse, random, showTag, showNumber, oneLine bool) {
	var stops map[int]string = map[int]string{
		1:  "[Часть] КПП",
		2:  "[Часть] Бухгалтерия",
		3:  "[Часть] Актовый зал",
		4:  "[Часть] Библиотека",
		5:  "[Часть] Медсанчасть",
		6:  "[Часть] Самоволка",
		7:  "[Часть] Дела",
		8:  "[Часть] Плац",
		9:  "[Часть] Холм",
		10: "[Часть] РЛС",

		11: "[КП] Вход в бункер",
		12: "[КП] Секретка",
		13: "[КП] Коммутатор",
		14: "[КП ЛАЗ] Дежурный",
		15: "[КП ЛАЗ] Громы",
		16: "[КП ЛАЗ] Койка",
		17: "[КП ЛАЗ] АТС",
		18: "[КП ЛАЗ] Ноутбук",
		19: "[КП ЛАЗ] Паяльник",
		20: "[КП ЛАЗ] Стол",

		21: "[КП Зал] Майор",
		22: "[КП Зал] Планшеты",
		23: "[КП ПРЦ] Приемники",
		24: "[КП ПРЦ] Новый год",
		25: "[КП] Телеграф",
		26: "[КП] Кладовка ", //the half of the deck
		27: "[Часть] Вышка",
		28: "[Часть] Релейка",
		29: "[Часть] Окоп",
		30: "[Часть] Автопарк",

		31: "[Казарма] Тумба",
		32: "[Казарма] Коптерка",
		33: "[Казарма] Сейф",
		34: "[Казарма] Штанги",
		35: "[Казарма] Кровать",
		36: "[Казарма] Бытовка",
		37: "[Казарма] Сушилка",
		38: "[Казарма] Канцелярия",
		39: "[Казарма] Комната досуга",
		40: "[Казарма] Оружейка",

		41: "[Часть] Курилка",
		42: "[Часть] Теплица",
		43: "[Часть] Спорт городок",
		44: "[Часть] Штаб",
		45: "[Вторая] Столовая",
		46: "[Вторая] Гостиница",
		47: "[Вторая] Турники",
		48: "[Вторая] Плац 2",
		49: "[Вторая] Склад",
		50: "[Вторая] Чепок",

		51: "[Дембель] Уазик",
		52: "[Дембель] Вокзал"}

	rx := regexp.MustCompile("\\[.*\\]\\s")

	fmt.Println()
	var k int
	for i := 1; i <= 52; i++ {
		k = i
		if reverse {
			k = 53 - i
		}

		//arrow
		if i != 1 {
			fmt.Print(" -> ")
			time.Sleep(time.Second * time.Duration(intervalInSeconds))
		}

		//tag & name
		tag := rx.FindString(stops[k])
		if showTag {
			fmt.Print(stops[k])
		} else {
			fmt.Print(strings.Replace(stops[k], tag, "", -1))
		}

		//number
		if showNumber {
			fmt.Print("(", k, ")")
			time.Sleep(time.Second)
		}

		if !oneLine {
			fmt.Println()
		}

		time.Sleep(time.Second)
		if oneLine {
			if !reverse && k%10 == 0 {
				fmt.Print("\n\n")
			} else if reverse && k%10 == 1 {
				fmt.Print("\n\n")
			}
		}
	}
}
