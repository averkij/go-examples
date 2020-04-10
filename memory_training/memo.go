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

	// showTags := true
	// sec := 2
	// itemsPerLine := 10
	// trainPersons(52, sec, itemsPerLine, showTags)

	//memory palace training

	reverse := false
	random := false
	showTag := false
	showNumber := true
	oneLine := true
	trainPalace(2, reverse, random, showTag, showNumber, oneLine)
}

func trainPersons(cardsAmount, intervalInSeconds, itemsPerLine int, showTags bool) {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	var cards map[int][]string = map[int][]string{
		0:  []string{"папаха    ", "b"},
		1:  []string{"ракета    ", "b"},
		2:  []string{"блюдо     ", "b"},
		3:  []string{"трость    ", "b"},
		4:  []string{"шарф      ", "b"},
		5:  []string{"протезы   ", "b"},
		6:  []string{"шпага     ", "b"},
		7:  []string{"коктейль  ", "b"},
		8:  []string{"монетка   ", "b"},
		9:  []string{"фут. мяч  ", "b"},
		10: []string{"зуб       ", "b"},
		11: []string{"робот     ", "b"},
		12: []string{"бас. мяч  ", "b"},
		13: []string{"часы      ", "b"},
		14: []string{"помада    ", "b"},
		15: []string{"перо      ", "b"},
		16: []string{"тату      ", "b"},
		17: []string{"леденец   ", "b"},
		18: []string{"пачка     ", "b"},
		19: []string{"листья    ", "b"},
		20: []string{"кресло    ", "b"},
		21: []string{"шест      ", "b"},
		22: []string{"сигарета  ", "b"},
		23: []string{"небоскрёб ", "b"},
		24: []string{"доспехи   ", "b"},
		25: []string{"нож       ", "b"},
		26: []string{"кросовки  ", "b"},
		27: []string{"гантеля   ", "b"},
		28: []string{"лопата    ", "b"},
		29: []string{"ударные   ", "b"},
		30: []string{"семечки   ", "b"},
		31: []string{"таблетки  ", "b"},
		32: []string{"молоток   ", "b"},
		33: []string{"винтовка  ", "b"},
		34: []string{"вилка     ", "b"},
		35: []string{"паяльник  ", "b"},
		36: []string{"молния    ", "b"},
		37: []string{"кисточка  ", "b"},
		38: []string{"г. клюшка ", "b"},
		39: []string{"виски     ", "b"},
		40: []string{"лапша     ", "b"},
		41: []string{"эл. гитара", "b"},
		42: []string{"лимон     ", "b"},
		43: []string{"бейсболка ", "b"},
		44: []string{"котелок   ", "b"},
		45: []string{"паутина   ", "b"},
		46: []string{"автомат   ", "b"},
		47: []string{"колонка   ", "b"},
		48: []string{"вагон м.  ", "b"},
		49: []string{"будда     ", "b"},
		50: []string{"конь      ", "b"},
		51: []string{"спас. круг", "b"},
		52: []string{"колода    ", "b"},
		53: []string{"костер    ", "b"},
		54: []string{"пиво      ", "b"},
		55: []string{"красн. куб", "b"},
		56: []string{"венок     ", "b"},
		57: []string{"бассейн   ", "b"},
		58: []string{"микрофон  ", "b"},
		59: []string{"пудель    ", "b"},
		60: []string{"скрипка   ", "b"},
		61: []string{"ребенок   ", "b"},
		62: []string{"кубик руб.", "b"},
		63: []string{"коньки    ", "b"},
		64: []string{"мобильник ", "b"},
		65: []string{"расчёска  ", "b"},
		66: []string{"горшок    ", "b"},
		67: []string{"лет. мышь ", "b"},
		68: []string{"скалка    ", "b"},
		69: []string{"диадема   ", "b"},
		70: []string{"ворота    ", "b"},
		71: []string{"скафандр  ", "b"},
		72: []string{"поб. слон ", "b"},
		73: []string{"танк      ", "b"},
		74: []string{"газета    ", "b"},
		75: []string{"волга     ", "b"},
		76: []string{"корова    ", "b"},
		77: []string{"радио     ", "b"},
		78: []string{"икебана   ", "b"},
		79: []string{"питон     ", "b"},
		80: []string{"приставка ", "b"},
		81: []string{"брусья    ", "b"},
		82: []string{"машина    ", "b"},
		83: []string{"груша     ", "b"},
		84: []string{"буденовка ", "b"},
		85: []string{"флаг      ", "b"},
		86: []string{"тайд      ", "b"},
		87: []string{"шахматы   ", "b"},
		88: []string{"гитара    ", "b"},
		89: []string{"рентген   ", "b"},
		90: []string{"обои      ", "b"},
		91: []string{"яхта      ", "b"},
		92: []string{"ж. человек", "b"},
		93: []string{"оборода   ", "b"},
		94: []string{"скальпель ", "b"},
		95: []string{"лыжи      ", "b"},
		96: []string{"кокос     ", "b"},
		97: []string{"доллары   ", "b"},
		98: []string{"метла     ", "b"},
		99: []string{"топор     ", "b"}}

	fmt.Println()

	for i := 0; i < cardsAmount; i++ {
		n := rnd.Int31n(100)
		if n < 10 {
			fmt.Print("0")
		}
		fmt.Print(n, " ")
		//time.Sleep(time.Second)

		if i != cardsAmount-1 {
			if ((i+1)%itemsPerLine != 0 && !showTags) || showTags {
				for j := 0; j < intervalInSeconds-1; j++ {
					fmt.Print("-")
					//time.Sleep(time.Second)
				}
				fmt.Print(" ")
			}
		}
		if showTags {
			fmt.Print(cards[int(n)][0], " | ")
			//time.Sleep(time.Second)
		}
		if (i+1)%itemsPerLine == 0 {
			fmt.Print("\n\n")
		}
	}
	fmt.Println()
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

	fmt.Print("\n")
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
		}
		time.Sleep(time.Second)

		if !oneLine {
			fmt.Println()
		} else {
			if !reverse && k%10 == 0 {
				fmt.Print("\n\n")
			} else if reverse && k%10 == 1 {
				fmt.Print("\n\n")
			}
		}
	}
}
