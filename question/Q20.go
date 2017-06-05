package main

import (
	"errors"
	"fmt"


)

func main() {
	str:=`{}()`
	//fmt.Println(reflect.TypeOf(str[1]))

	fmt.Println(isValid(str))
	//fmt.Println(string(str[1]))
}

func isValid(s string) bool {
	if len(s)==0{
		return false
	}
	var stack Stack
	var i int
	for i=0;i<len(s);i++{
		switch s[i] {
		case 40:
			stack.Push(s[i])
		case 91:
			stack.Push(s[i])
		case 123:
			stack.Push(s[i])
		case 41:
			stackPop,_:=stack.Pop()
			if stackPop!=40 {
				return false
			}
		case 93:
			stackPop,_:=stack.Pop()
			if stackPop!=91 {
				return false
			}
		case 125:
			stackPop,_:=stack.Pop()
			if stackPop!=123 {
				fmt.Println(stackPop,"error")
				return false
			}
		default:
			return false
		}
	}
	if stack.Len()==0{
		return true
	}
	return false
}
type Stack []uint8

func(stack *Stack) Push(v uint8) {
	*stack = append(*stack, v)
}

func(stack *Stack) Pop() (uint8, error){
	if len(*stack) == 0 {
		return 0, errors.New("stack empty")
	}
	v := (*stack)[len(*stack) - 1]
	*stack = (*stack)[:len(*stack) - 1]
	return v, nil
}

func(stack *Stack) Top() (uint8, error) {
	if len(*stack) == 0 {
		return 0, errors.New("stack empty")
	}

	return (*stack)[len(*stack) - 1], nil
}

func(stack *Stack) Len() int {
	return len(*stack)
}