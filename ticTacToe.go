package main

import "fmt"

var tab = [3][3]string{{"x", "o", "x"}, {"o", "o", "x"}, {"_", "_", "_"}}

// x - computer || o - human
// -1 - champion=human || 0 - nichaya || +1 - champion-computer

// clean tsble
//var tab = [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}

//var numeric_coord = [9][2]int{{0,0},{0,1},{0,2},{1,0},{1,1},{1,2},{2,0},{2,1},{2,2}}

/*
_x_
oxx
_oo

xx_
oxx
_oo

_xx
oxx
_oo

_x_
oxx
xoo
*/

func printTab(tab [3][3]string) {
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			fmt.Printf(tab[y][x] + " | ")
		}
		fmt.Println()
	}
}

func isEndOfTheGame(tab [3][3]string) bool {

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if tab[y][x] == "_" {
				return false
			}
		}
	}
	return true
}

func isChampion(tab [3][3]string) string {

	//horizont champion
	for y := 0; y < 3; y++ {
		if tab[y][0] == tab[y][1] && tab[y][1] == tab[y][2] && tab[y][0] != "_" {
			return tab[y][0]
		}
	}

	// vertical champion
	for x := 0; x < 3; x++ {
		if tab[0][x] == tab[1][x] && tab[1][x] == tab[2][x] && tab[0][x] != "_" {
			return tab[0][x]
		}
	}

	//diagonal champion
	if (tab[0][0] == tab[1][1] && tab[1][1] == tab[2][2] && tab[0][0] != "_") ||
		(tab[0][2] == tab[1][1] && tab[1][1] == tab[2][0] && tab[1][1] != "_") {
		return tab[1][1]
	}

	// нет выйгрывших
	return "_"
}

func printArrTab(arrTab [][3][3]string) {
	for i := 0; i < len(arrTab); i++ {
		printTab(arrTab[i])
		fmt.Println()
	}
}

func getArrTab(tab [3][3]string, gamer string) [][3][3]string {
	var arrTab [][3][3]string
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if tab[y][x] == "_" {
				newTab := tab
				newTab[y][x] = gamer
				arrTab = append(arrTab, newTab)
			}
		}
	}
	return arrTab
}

// минимах алгоритм
// функции высшего порядка

// ДЗ
//интерфейс чтобы играть с самим собой
//пронумеровать клетки

func changeGamer(gamer string) string {
	if gamer == "o" {
		return "x"
	} else {
		return "o"
	}
}

func count_variants(tab [3][3]string, gamer string) int {
	printTab(tab)
	fmt.Println()

	if isEndOfTheGame(tab) {
		return 1
	}

	champion := isChampion(tab)
	if champion == "x" || champion == "o" {
		return 1
	}

	arrOfTab := getArrTab(tab, gamer)
	countNum := 0

	newGamer := changeGamer(gamer)
	for i := 0; i < len(arrOfTab); i++ {
		countNum += count_variants(arrOfTab[i], newGamer)
	}
	return countNum + 1
}

func getMax(arr []int8) int8 {
	maxVal := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}

	}
	return maxVal
}

func getMin(arr []int8) int8 {
	minVal := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < minVal {
			minVal = arr[i]
		}
	}
	return minVal
}

func checkTable(tab [3][3]string, gamer string) int8 {
	champion := isChampion(tab)
	if champion == "x" {
		return 1
	} else if champion == "o" {
		return -1
	} else if isEndOfTheGame(tab) {
		return 0
	}

	arrOfChecks := []int8{}
	tabVariants := getArrTab(tab)
	for i := 0; i < len(tabVariants); i++ {
		arrOfChecks = append(arrOfChecks, checkTable(tabVariants[i], gamer))
	}

	if gamer == "x" {
		return getMax(arrOfChecks)
	} else {
		return getMin(arrOfChecks)
	}
}

//func make_move(tab [3][3]string, x int, y int, gamer string) [3][3]string {
//
//	return tab
//}u

func start_the_game(tab [3][3]string, gamer string) string {
	fmt.Println("*** Игра началась ***")
	championLetter := ""

	for i := 1; i <= 9; i++ {
		var x, y int
		fmt.Println("--- Ход № ", i, " Игрок: ", gamer)
		printTab(tab)
		fmt.Println("Ваш ход (столбец):")
		fmt.Scanf("%d\n", &x)
		fmt.Println("Ваш ход (строка):")
		fmt.Scanf("%d\n", &y)
		tab[y][x] = gamer
		if isChampion(tab) != "_" {
			championLetter = "!!!!! Победили " + gamer + " !!!!!!"
			break
		}
		if isEndOfTheGame(tab) {
			championLetter = "!!!!! Ничья !!!!!"
			break
		}

		gamer = changeGamer(gamer)
		fmt.Println("=================================")
		fmt.Println()
	}
	return championLetter
}

func main() {
	//print_tab(tab)
	//fmt.Println(is_end_of_the_game(tab))
	//arrTab := get_arr_tab(tab, "o")
	//printArrTab(arrTab)
	//fmt.Println(count_variants(tab, "o"))
	//fmt.Println(start_the_game(tab, "x"))
	fmt.Println(checkTable(tab, "x"))
}
