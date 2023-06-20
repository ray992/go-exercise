package main

import (
	"fmt"
	"strconv"
)

/**
	逆波兰式
**/

func evalRPN(tokens []string) int {
	var numSlice []int
	n := len(tokens)
	for i := 0; i < n; i++ {
		curToken := tokens[i]
		if curToken == "+" || curToken == "-" || curToken == "*" || curToken == "/" {
			pre2 := numSlice[len(numSlice) - 2]
			pre1 := numSlice[len(numSlice) - 1]
			numSlice = numSlice[0:len(numSlice) - 2]
			if curToken == "+" {
				numSlice = append(numSlice, pre1 + pre2)
			}else if curToken == "-" {
				numSlice = append(numSlice, pre2 - pre1)
			}else if curToken == "*" {
				numSlice = append(numSlice, pre1 * pre2)
			}else {
				numSlice = append(numSlice, pre2 / pre1)
			}
 		} else {
			v, _ := strconv.Atoi(curToken)
			numSlice = append(numSlice, v)
		}
	}
	return numSlice[len(numSlice) - 1]
}

// 改用switch

func evalRPN2(tokens []string) int {
	var numSlice []int
	n := len(tokens)
	for i := 0; i < n; i++ {
		curToken := tokens[i]
		v, err := strconv.Atoi(curToken)
		if err == nil {
			numSlice = append(numSlice, v)
		}else {
			pre2 := numSlice[len(numSlice) - 2]
			pre1 := numSlice[len(numSlice) - 1]
			numSlice = numSlice[0:len(numSlice) - 2]
			switch curToken{
			case "+":
				numSlice = append(numSlice, pre1 + pre2)
			case "-":
				numSlice = append(numSlice, pre2 - pre1)
			case "*":
				numSlice = append(numSlice, pre1 * pre2)
			case "/":
				numSlice = append(numSlice, pre2 / pre1)
			}
		}
	}
	return numSlice[len(numSlice) - 1]
}




func main() {
	fmt.Println(evalRPN2([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN2([]string{"4","13","5","/","+"}))
	fmt.Println(evalRPN2([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))
}