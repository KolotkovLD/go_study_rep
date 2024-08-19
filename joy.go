// 1 + 1
// 1 1 +
// 4 3 *
// 4 * 3 

// 2 2 +  5 5 /  * stack <- <-  (с вершины stack)
//( 2+2)*(5/5)



//токенизация - разбивка строки по пробелу

//дз - представление гетерогенных массивов (числа строки и другие массивы)
// array_intedrace := []interface{}{}


package main

import (
    "fmt"
    "strconv"
)


func tokenizer (str string) []string{
    token_list := []string{}
    str_item := ""
    for i:=0 ; i<len(str); i++{
	if str[i] == ' '{
	    if str_item != ""{
		token_list = append(token_list, str_item)
	    }
	    str_item = ""
	}else{
	    str_item = str_item + string(str[i])
	}
    }
    
    if len(str_item)!=0{
        token_list = append(token_list, str_item)
    }
    return token_list
}

func operation(operandF int, operandS int, operation string) int{
    result := 0
    switch operation{
	case "+": result = operandF + operandS
	case "-": result = operandF - operndS
	case "*": result = operandF * operandS
	case "/": result = operandF / operandS
    }
    return result
}

func calculator(list []string) int{
    stack := []int
    for i:= ; i<len(list); i++{
	if num, err: strconv.Atoi(list[i]); err == nil{
	    stack = append(stack, num)
	}else{
	    operandF := stack[len(stack)-2]
	    operandS := stack[len(stack)-1]
	    result := operation(operandF, operandS, list[i])
	    stack = stack[:len(stack)-2]
	    stack = append(stack, result)
	}
    }
    return stack[0]
}


func main(){
    //fmt.Println(tokenizer(
    oper := "2 122  + 5 5 / *"
    fmt.Println(calculator(tokenizer(oper)))
}

