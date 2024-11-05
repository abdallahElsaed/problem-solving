package main

import (
	// "strings"
	// "unicode"
	"strconv"
)
func evalRPN(tokens []string) int {
    
	// loop over the array 
	// make condition :
	// if the index value is number push in stack array
	// if  not: 
		// get the top element and assign to n2 and remove from stack
		// get the top element again and assign to n1 and remove from stack
		// calculate n1 (the current element in loop is the operator) n2 [ex: 2(n1) + 1(n1)]
		// push the result in stack again
	stack := []int{}
	for _, value := range tokens{
		
		result :=0
		switch value {
		case "+":
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = n1 + n2
			stack = append(stack, result)
		case "-":
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = n1 - n2
			stack = append(stack, result)
		case "*":
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = n1 * n2
			stack = append(stack, result)
		case "/":
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = n1 / n2
			stack = append(stack, result)
		default:
			v, _:= strconv.Atoi(value)
			stack = append(stack, v)
		}

	}
	return stack[0]
}




/*



// if unicode.IsNumber(rune(value[0])) {
		// 	v, _:= strconv.Atoi(value)
		// 	fmt.Println("if value",v)

		// 	stack = append(stack, v)
		// 	fmt.Println("if number",stack)
		// }else{
		// 	fmt.Println("else value",value)

		// 	n2 := stack[len(stack)-1]
		// 	stack = stack[:len(stack)-1]
		// 	n1 := stack[len(stack)-1]
		// 	stack = stack[:len(stack)-1]
		// 	result :=0
		// 	switch value {
		// 	case "+":
		// 		result = n1 + n2
		// 		fmt.Println("cas +",result)
		// 	case "-":
		// 		result = n1 - n2
		// 	case "*":
		// 		result = n1 * n2
		// 	case "/":
		// 	result = n1 / n2
		// 	}
		// 	fmt.Println("result",result)
		// 	stack = append(stack, result)
		// 	fmt.Println("after",stack)

		// }
*/