package main
import "fmt"

var tab = [3][3]string{{"o","o","x"},{"x","_","o"},{"o","_","x"}}

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

func getChampion(tab [3][3]string) string{
    //ДЗ
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

func main(){
    //printTab(tab)
    //fmt.Println(isEndOfTheGame(tab))
    arrTab := getArrTab(tab, "o")
    printArrTab(arrTab)
}