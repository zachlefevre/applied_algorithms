package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	for ind := 0; ind < len(tokens); ind++ {
		token := tokens[ind]
		fmt.Println("token: ", token, "tokens: ", tokens, "index: ", ind, "len: ", len(tokens))
		var el int
		switch c := token; c {
		case "*":
			l, _ := strconv.Atoi(tokens[ind-2])
			r, _ := strconv.Atoi(tokens[ind-1])
			el = l * r
			lArr := tokens[:ind-2]
			rArr := tokens[ind+1:]
			fmt.Println(lArr, rArr)
			tokens = append(lArr, append([]string{strconv.Itoa(el)}, rArr...)...)
			ind -= 2
		case "/":
			l, _ := strconv.Atoi(tokens[ind-2])
			r, _ := strconv.Atoi(tokens[ind-1])
			el = l / r
			lArr := tokens[:ind-2]
			rArr := tokens[ind+1:]
			fmt.Println(lArr, rArr)
			tokens = append(lArr, append([]string{strconv.Itoa(el)}, rArr...)...)
			ind -= 2
		case "+":
			l, _ := strconv.Atoi(tokens[ind-2])
			r, _ := strconv.Atoi(tokens[ind-1])
			el = l + r
			lArr := tokens[:ind-2]
			rArr := tokens[ind+1:]
			fmt.Println(lArr, rArr)
			tokens = append(lArr, append([]string{strconv.Itoa(el)}, rArr...)...)
			ind -= 2
		case "-":
			l, _ := strconv.Atoi(tokens[ind-2])
			r, _ := strconv.Atoi(tokens[ind-1])
			el = l - r
			lArr := tokens[:ind-2]
			rArr := tokens[ind+1:]
			fmt.Println(lArr, rArr)
			tokens = append(lArr, append([]string{strconv.Itoa(el)}, rArr...)...)
			ind -= 2
		default:
		}
	}
	ret, _ := strconv.Atoi(tokens[0])
	return ret
}
func main() {
	fmt.Println("val: ", evalRPN([]string{"4", "2", "8", "*", "*", "3", "+"}))
}
