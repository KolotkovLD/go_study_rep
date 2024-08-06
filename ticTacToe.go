package main
import "fmt"

var tab = [3][3]string{{"x","o","x"},{"o","o","x"},{"_","_","_"}}
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

func printTab(tab [3][3]string){
    for y:=0; y<3; y++{
	for x:=0; x<3; x++{
	    fmt.Printf(tab[y][x] + " | ")
	}
	fmt.Println()
    }
}



func isEndOfTheGame(tab [3][3]string) bool{
    
    for y:=0; y<3; y++{
	for x:=0; x<3; x++{
	    if (tab[y][x] == "_"){
		return  false
	    }
	}
    }
    return true
}

func getChampion(tab [3][3]string, gamer string) string{
    //horizont champion
    for y:=0; y<3; y++{
	if gamer == tab[y][0] && gamer == tab[y][1] && gamer == tab[y][2]{
	    return 
	}
    }
    return ""
}

func printArrTab(arrTab [][3][3]string){
    for i:=0; i<len(arrTab); i++{
	printTab(arrTab[i])
	fmt.Println()
    }
}

func getArrTab(tab [3][3]string, gamer string) [][3][3]string{
    var arrTab [][3][3]string
    for y:=0; y<3; y++{
	for x:=0; x<3; x++{
	    if tab[y][x] == "_"{
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

func get_gamer(gamer string) string{
    if gamer == "o"{
    	gamer = "x"
    }else{
	gamer = "o"
    }
    return gamer
}

func count_variants(tab [3][3]string, gamer string) int{
    printTab(tab)
    fmt.Println()

    new_gamer := get_gamer(gamer)
    arr_of_tab := getArrTab(tab, new_gamer)
    count_num := 0

    if len(arr_of_tab) == 1{
	return 1
    }

    for i:=0; i<len(arr_of_tab);i++{
	count_num += count_variants(arr_of_tab[i], new_gamer)
    }
    return count_num + 1
}

func main(){
    //printTab(tab)
    //fmt.Println(isEndOfTheGame(tab))
    //arrTab := getArrTab(tab, "o")
    //printArrTab(arrTab)
    fmt.Println(count_variants(tab, "o"))
}