package main
import "fmt"

func esc_sort(arr []int) bool{
    for i :=1 ; i < len(arr); i++{
	if (arr[i-1] > arr[i]){
	    return false
	}
    }
    return true
}


func bubble_sort(arr []int) []int{
    for (!esc_sort(arr)){
	for i := 1; i < len(arr); i++{
	    if (arr[i-1] > arr[i]){
		prv_item := arr[i-1]
		arr[i-1] = arr[i]
		arr[i] = prv_item
	    }
	}
    }
    return arr
}

func max_val(arr []int) int{
    max_item := arr[0]
    for i := 1; i < len(arr); i++{
	if (arr[i] > max_item){
	    max_item = arr[i]
	}
    }
    return max_item
}

1 
3
4 
5
6 
8
//поиск минимального, с дальнейшим поиском в усеченном массиве
//терминал линукса
//git


3 4 5 6 8 9
//перестановка парой (идем слева направо, сравниваем, отскакиваем, если корректно идем вправо)

func main(){
    fmt.Println(esc_sort([]int{2,4,6,3}))
    fmt.Println(esc_sort([]int{1,2,3,4,5}))

    fmt.Println(bubble_sort([]int{1,3,2,4, 66, 76, 5}))
    fmt.Println(max_val([]int{1, 4, 66, 7}))
}