package main

import "fmt"

func esc_sort(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func bubble_sort(arr []int) []int {
	for !esc_sort(arr) {
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				prv_item := arr[i-1]
				arr[i-1] = arr[i]
				arr[i] = prv_item
			}
		}
	}
	return arr
}

func max_val(arr []int) int {
	max_item := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max_item {
			max_item = arr[i]
		}
	}
	return max_item
}

//1
//3/
//4
//5
//6
//8
//поиск минимального, с дальнейшим поиском в усеченном массиве
//терминал линукса
//git

//3 4 5 6 8 9
//перестановка парой (идем слева направо, сравниваем, отскакиваем, если корректно идем вправо)

//{3, 6, 2, 7, 78, 0, 62, 12, 9}
//2, 0   ->  0, 2
//6,7, 78, 62, 12, 9   ->   6, 7, 9, 12, 62, 78

//0, 2, 3  , 6, 7, 9 ,12, 62, 78

func pare_sort(arr []int) []int {
	for i := 1; i < len(arr); {
		if arr[i-1] > arr[i] {
			prv_item := arr[i-1]
			arr[i-1] = arr[i]
			arr[i] = prv_item
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}
	return arr
}

func part_sort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	var new_arr_more []int
	var new_arr_less []int

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[0] {
			new_arr_more = append(new_arr_more, arr[i])
		} else {
			new_arr_less = append(new_arr_less, arr[i])
		}
	}

	new_arr_more = part_sort(new_arr_more)
	new_arr_less = part_sort(new_arr_less)
	new_arr_less = append(new_arr_less, arr[0])
	new_arr_less = append(new_arr_less, new_arr_more...)
	return new_arr_less
}

/*
6, 4, 2, 7, 8, 5
6, 4, 2    7, 8, 5
, ,     , 7, 8

2, 4, 5, 6, 7, 8


*/

func middle_sort() {}

func compare_sort(arr1 []int, arr2 []int) []int {
	var arr_rt = []int{}

	for len(arr1) > 0 || len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			arr_rt = append(arr_rt, arr1[0])
			arr1 = arr1[1:]
		} else {
			arr_rt = append(arr_rt, arr2[0])
			arr2 = arr2[1:]
		}

	}
	return arr_rt
}

func main() {
	fmt.Println(esc_sort([]int{2, 4, 6, 3}))
	fmt.Println(esc_sort([]int{1, 2, 3, 4, 5}))

	fmt.Println(bubble_sort([]int{1, 3, 2, 4, 66, 76, 5}))
	fmt.Println(max_val([]int{1, 4, 66, 7}))
	fmt.Println(part_sort([]int{0, 2, 9, 6, 7, 78, 62}))
	fmt.Println()
	fmt.Println(pare_sort([]int{3, 2, 5, 6, 0, 9}))
}
