package main

import "fmt"

//var tab = [3][3]string{{"x","o","x"},{"o","o","x"},{"_","_","_"}}

// clean tsble
var tab = [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}

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

func print_tab(tab [3][3]string) {
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			fmt.Printf(tab[y][x] + " | ")
		}
		fmt.Println()
	}
}

func is_end_of_the_game(tab [3][3]string) bool {

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if tab[y][x] == "_" {
				return false
			}
		}
	}
	return true
}

func is_champion(tab [3][3]string, gamer string) bool {
	//horizont champion
	for y := 0; y < 3; y++ {
		if gamer == tab[y][0] && gamer == tab[y][1] && gamer == tab[y][2] {
			return true
		}
	}

	// vertical champion
	for x := 0; x < 3; x++ {
		if gamer == tab[0][x] && gamer == tab[1][x] && gamer == tab[2][x] {
			return true
		}
	}

	//diagonal champion
	if (gamer == tab[0][0] && gamer == tab[1][1] && gamer == tab[2][2]) ||
		(gamer == tab[0][2] && gamer == tab[1][1] && gamer == tab[2][0]) {
		return true
	}

	return false
}

func printArrTab(arrTab [][3][3]string) {
	for i := 0; i < len(arrTab); i++ {
		print_tab(arrTab[i])
		fmt.Println()
	}
}

func get_arr_tab(tab [3][3]string, gamer string) [][3][3]string {
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

// ДЗ
//интерфейс чтобы играть с самим собой
//пронумеровать клетки

func change_gamer(gamer string) string {
	if gamer == "o" {
		gamer = "x"
	} else {
		gamer = "o"
	}
	return gamer
}

func count_variants(tab [3][3]string, gamer string) int {
	print_tab(tab)
	fmt.Println()

	if is_end_of_the_game(tab) {
		return 1
	}

	if is_champion(tab, gamer) {
		return 1
	}

	new_gamer := change_gamer(gamer)
	arr_of_tab := get_arr_tab(tab, new_gamer)
	count_num := 0

	if len(arr_of_tab) == 1 {
		return 1
	}

	for i := 0; i < len(arr_of_tab); i++ {
		count_num += count_variants(arr_of_tab[i], new_gamer)
	}
	return count_num + 1
}

//func make_move(tab [3][3]string, x int, y int, gamer string) [3][3]string {
//
//	return tab
//}


func start_the_game(tab [3][3]string, gamer string) string {
	fmt.Println("*** Игра началась ***")
	champion_letter := ""

	for i := 1; i<=9; i++ {
		var x, y int
		fmt.Println("--- Ход № ", i, " Игрок: ", gamer)
		print_tab(tab)
		fmt.Println("Ваш ход (столбец):")
		fmt.Scanf("%d\n", &x)
		fmt.Println("Ваш ход (строка):")
		fmt.Scanf("%d\n", &y)
		tab[y][x] = gamer
		if is_champion(tab, gamer){
			champion_letter = "!!!!! Победили " + gamer + " !!!!!!"
			break
		}
		if is_end_of_the_game(tab) {
			champion_letter = "!!!!! Ничья !!!!!"
			break
		}

		gamer = change_gamer(gamer)
		fmt.Println("=================================")
		fmt.Println()
	}
	return champion_letter
}

func main() {
	//print_tab(tab)
	//fmt.Println(is_end_of_the_game(tab))
	//arrTab := get_arr_tab(tab, "o")
	//printArrTab(arrTab)
	//fmt.Println(count_variants(tab, "o"))
	fmt.Println(start_the_game(tab, "x"))
}
