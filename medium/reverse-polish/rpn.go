package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"strconv"
)

func evalRPN(tokens []string) int {
	stk := stack.New()
	fmt.Println(stk)

	for _, token := range tokens {
		switch c := token; c {
		case "*":
			fmt.Println(c)
			stk.Push(int(stk.Pop().(int) * stk.Pop().(int)))
		case "/":
			fmt.Println(c)
			stk.Push(stk.Pop().(int) / stk.Pop().(int))
		case "+":
			fmt.Println(c)
			stk.Push(stk.Pop().(int) + stk.Pop().(int))
		case "-":
			fmt.Println(c)
			stk.Push(stk.Pop().(int) - stk.Pop().(int))
		default:
			v, _ := strconv.Atoi(token)
			fmt.Println("pushing", v)
			stk.Push(v)
		}
	}
	return stk.Pop().(int)
}

func main() {
	fmt.Println(evalRPN([]string{"3", "2", "4", "+", "/", "5", "*"}))
}
