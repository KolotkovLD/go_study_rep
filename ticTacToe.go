package main

import "fmt"

var tab = [3][3]string{{"x","o","x"},{"o","o","x"},{"_","_","_"}}
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

func is_champion(tab [3][3]string) string {
	
	//horizont champion
	for y := 0; y < 3; y++ {
		if tab[y][0] == tab[y][1] && tab[y][1] == tab[y][2] && tab[y][0] != "_" {
			return tab[y][0]
		}
	}

	// vertical champion
	for x := 0; x < 3; x++ {
		if tab[0][x] == tab[1][x] && tab[1][x] == tab[2][x] && tab[0][x] != "_"{
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


// минимах алгоритм
// функции высшего порядка
 


// ДЗ
//интерфейс чтобы играть с самим собой
//пронумеровать клетки

func change_gamer(gamer string) string {
	if gamer == "o" {
		return "x"
	} else {
		return "o"
	}
}

func count_variants(tab [3][3]string, gamer string) int {
	print_tab(tab)
	fmt.Println()

	if is_end_of_the_game(tab) {
		return 1
	}

	champion := is_champion(tab)
	if champion == "x" || champion == "o" {
		return 1
	}

	arr_of_tab := get_arr_tab(tab, gamer)
	count_num := 0

	new_gamer := change_gamer(gamer)
	for i := 0; i < len(arr_of_tab); i++ {
		count_num += count_variants(arr_of_tab[i], new_gamer)
	}
	return count_num + 1
}



func get_max(arr []int8)int8{
    max_val := arr[0]
    for i:=1; i<len(arr); i++{
	if arr[i] > max_val{
	    max_val =  arr[i]
	}
	
    }
    return max_val
}

func get_min(arr []int8)int8{
    min_val := arr[0]
    for i:=0; i<len(arr); i++{
	if arr[i] < min_val{
	    min_val = arr[i]
	}
    }
    return min_val
}

func check_table(tab [3][3]string, gamer string) int8{
    champion := is_champion(tab)
    if champion == "x"{
	return 1
    }else if champion == "o"{
	return -1
    }else if is_end_of_the_game(tab){
	return 0
    }
    
    arr_of_checks := []int8{}
    tab_variants := get_arr_tab(tab)
    for i:=0; i<len(tab_variants); i++{
        arr_of_checks = append(arr_of_checks, check_table(tab_variants[i], gamer))
    }

    if gamer == "x"{
	return get_max(arr_of_checks)
    }else{
	return get_min(arr_of_checks)
    }
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
		if is_champion(tab) != "_"{
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
	//fmt.Println(start_the_game(tab, "x"))
	fmt.Println(check_table(tab, "x"))
}
