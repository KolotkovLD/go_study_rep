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
6, 4, 2, |  7, 8, 5
   |   7, 8

2, 4, 5, 6, 7, 8

*/

func quick_sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		pivot := arr[len(arr)/2]
		lesser := []int{}
		for i := 0; len(arr) > i; i++ {
			if pivot > arr[i] {
				lesser = append(lesser, arr[i])
			}
		}
		bigger := []int{}
		for i := 0; len(arr) > i; i++ {
			if pivot < arr[i] {
				bigger = append(bigger, arr[i])
			}
		}
		rt_arr := []int{}
		rt_arr = append(rt_arr, quick_sort(lesser)...)
		rt_arr = append(rt_arr, pivot)
		rt_arr = append(rt_arr, quick_sort(bigger)...)
		return rt_arr
	}
}

func merge_sort(arr []int) []int {
	if len(arr) == 1 || len(arr) == 0 {
		return arr
	}
	arr1 := merge_sort(arr[:len(arr)/2])
	arr2 := merge_sort(arr[len(arr)/2:])
	return compare_sort(arr1, arr2)
}

func compare_sort(arr1 []int, arr2 []int) []int {
	var arr_rt = []int{}

	for len(arr1) > 0 && len(arr2) > 0 {
		//fmt.Println(arr1[0], arr2[0])
		if arr1[0] < arr2[0] {
			arr_rt = append(arr_rt, arr1[0])
			arr1 = arr1[1:]
		} else {
			arr_rt = append(arr_rt, arr2[0])
			arr2 = arr2[1:]
		}

	}
	if len(arr1) == 0 {
		arr_rt = append(arr_rt, arr2...)
	}
	if len(arr2) == 0 {
		arr_rt = append(arr_rt, arr1...)
	}
	return arr_rt
}

/*число четное если предыдущее нечетное
число нечетное если предыдущее четное*/

func even_num(num int) bool {
	if num == 0 {
		return true
	}
	if odd_num(num - 1) {
		return true
	}
	return false
}

func odd_num(num int) bool {
	if num == 0 {
		return false
	}
	if even_num(num - 1) {
		return true
	}
	return false
}

/*найти рекурсивно максимальный элемент массива*/

func requr_max(arr []int) int {

	if len(arr) == 1 {
		return arr[0]
	}
	req_val := requr_max(arr[1:])
	if arr[0] > req_val {
		return arr[0]
	} else {
		return req_val
	}
}

func reverse_string(str string) string {
	if len(str) == 0 {
		return ""
		//}
		//if len(str) == 1 {
		//	return string(str)
	} else {
		firstChar := str[0]
		restOfChar := str[1:]
		reversedRest := reverse_string(restOfChar)
		return reversedRest + string(firstChar)
	}
}

/*    rivet P
ivet	r	P
vet		i	r	P
et		v	irP
t
*/

func reqursion_append(arr1 []int, arr2 []int) []int {
	if len(arr2) == 1 {
		return append(arr1, arr2[0])
	}
	return reqursion_append(append(arr1, arr2[0]), arr2[1:])
}

/*
1) рекурсивная функция, которая переворачивает строку
2) рекурсивное исполнение функции которая соединяет массивы   compare_sort
3) --подготовка к "крестикам ноликам"
*/

// рекурсивная функция по получению переворотов 2024_0_30
//abc		acb	acb	cba	cab	abc	bca
// bc 	cb	bc	\	acb	abc
// ac	ca	ac	\	bca	bac
// ab	ba	ab	\	cab	cba
// acb abc bca bac cab cba
//

func req_per(str string) []string {
	/*first_str := str[0]
	  part1_str := str[1:]
	  part2_str := str[0]+str[2:]
	  part3_str := str[0:2] +str[3:]*/

	fmt.Println(str)
	if len(str) == 1 {
		return []string{str}
	}
	rt_arr := []string{}
	for i := 0; len(str) > i; i++ {
		var new_item string
		new_item = str[:i] + str[i+1:]
		per_arr := req_per(new_item)
		for j := 0; len(per_arr) > j; j++ {
			rt_arr = append(rt_arr, string(str[i])+string(per_arr[j]))
		}

	}
	return rt_arr
}

//https://www.mathsisfun.com/games/towerofhanoi.html
//рекурсивная пирамида
func piramid(start int, middle int, fin int, count int) {
	if count == 0 {
		fmt.Println("Конец")
	} else {
		piramid(start, fin, middle, count-1)
		fmt.Println("Переносим большой на последнюю палку: со ", start, " на ", fin)
		piramid(middle, start, fin, count-1)
	}

}

func main() {
	/*fmt.Println(esc_sort([]int{2, 4, 6, 3}))
	fmt.Println(esc_sort([]int{1, 2, 3, 4, 5}))

	fmt.Println(bubble_sort([]int{1, 3, 2, 4, 66, 76, 5}))
	fmt.Println(max_val([]int{1, 4, 66, 7}))
	fmt.Println(part_sort([]int{0, 2, 9, 6, 7, 78, 62}))
	fmt.Println()
	fmt.Println(pare_sort([]int{3, 2, 5, 6, 0, 9}))
	fmt.Println(quick_sort([]int{3, 2, 5, 6, 0, 9}))

	fmt.Println(compare_sort([]int{2, 5, 88, 100}, []int{7, 44, 101}))
	fmt.Println(merge_sort([]int{88, 5, 2, 100, 7, 101, 44}))
	fmt.Println(even_num(101))
	fmt.Println(odd_num(101))
	fmt.Println()
	//fmt.Println(requr_max([]int{}))*/
	//fmt.Println(reverse_string(""))
	//fmt.Println(reqursion_append([]int{1, 2, 3}, []int{5, 6, 7}))
	//fmt.Println(req_per("abc"))
	piramid(1, 2, 3, 3)
}
