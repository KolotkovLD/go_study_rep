// 1 + 1
// 1 1 +
// 4 3 *
// 4 * 3

// 2 2 +  5 5 /  * stack <- <-  (с вершины stack)
//( 2+2)*(5/5)

//токенизация - разбивка строки по пробелу

//дз - представление гетерогенных массивов (числа строки и другие массивы)
// array_interface := []interface{}{}
// append
// interface
//

package main

import (
	"fmt"
	"strconv"
)

func tokenizer(str string) []string {
	token_list := []string{}
	str_item := ""
	for i := 0; i < len(str); i++ {
		if str[i] == '[' || str[i] == ']' {
			if len(str_item) != 0 {
				token_list = append(token_list, str_item)
			}
			token_list = append(token_list, string(str[i]))
			str_item = ""
		} else if str[i] == ' ' {
			if str_item != "" {
				token_list = append(token_list, str_item)
			}
			str_item = ""
		} else {
			str_item = str_item + string(str[i])
		}
	}

	if len(str_item) != 0 {
		token_list = append(token_list, str_item)
	}
	return token_list
}

func operation(operandF int, operandS int, operation string) int {
	result := 0
	switch operation {
	case "+":
		result = operandF + operandS
	case "-":
		result = operandF - operandS
	case "*":
		result = operandF * operandS
	case "/":
		result = operandF / operandS
	}
	return result
}

func calculator(list []string) int {
	stack := []int{}
	for i := 0; i < len(list); i++ {
		if num, err := strconv.Atoi(list[i]); err == nil {
			stack = append(stack, num)
		} else {
			operandF := stack[len(stack)-2]
			operandS := stack[len(stack)-1]
			result := operation(operandF, operandS, list[i])
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		}
	}
	return stack[0]
}

// функция массив токенов - все скобки имеют пары -> bool
// dddddddd[ dsddddd [d dd    [rfrf frrr]]]

// [1 2 3]first 4 +

func checkBrackets(arr []string) bool {
	haveOpenBracket := false
	openBrackets := 0  // Счетчик открытых скобок
	closeBrackets := 0 // Счетчик закрытых скобок
	var checkInside []bool
	var finCheckInside = true
	for ind, char := range arr {
		switch char {
		case "[":
			if haveOpenBracket == false {
				openBrackets++ // Открывающая скобка увеличивает счетчик открытых скобок
			} else {
				checkInside = append(checkInside, checkBrackets(arr[ind:]))
			}
		case "]":
			closeBrackets++ // Закрывающая скобка увеличивает счетчик закрытых скобок
		}
	}
	for _, chck := range checkInside {
		if chck == false {
			finCheckInside = false
		}
	}
	if openBrackets == closeBrackets && finCheckInside == true {
		return true
	}
	return false
}

func main() {
	//fmt.Println(tokenizer(
	oper := "[2]first[122 [ vvv]] +  5 5 / *"
	token_list := tokenizer(oper)
	//fmt.Println(len(token_list))
	for _, v := range token_list {
		fmt.Println(v)
	}
	fmt.Println(checkBrackets(token_list))
	//fmt.Println(calculator(tokenizer(oper)))
	//arr := []interface{}{1, 2, "str"}
	//intF, _ := arr[0].(int)
	//intS, _ :=  arr[1].(int)
	//res := intF + intS
	//fmt.Println(res)
}
