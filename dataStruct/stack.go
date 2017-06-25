package main

import "fmt"

type elemType int

type Stack struct {
	Val elemType
	Next *Stack
}

func (S *Stack)Push(value elemType){
	fmt.Println("s->",S)
	newNode:=Stack{value,nil}
	if S == nil {
		fmt.Println("我为空")
		*S = newNode
	} else {
		newNode.Next=S
		S=&newNode
		fmt.Println("push->",)
	}

}


func (S *Stack)Pop() elemType{
	if S==nil {
		fmt.Println("this stack is nil!")
		return -1
	}
	elem:=S.Val
	*S=*(S.Next)
	return elem
}

func IsNil(s *Stack) bool {
	return s==nil
}
func PrintAll(s *Stack)  {
	if s==nil{
		return
	}
	for !IsNil(s){
		//fmt.Println(Pop(s))
	}
}

func main()  {
	s:=new(Stack)
	s.Val=1
	s.Push(2)
	s.Push(3)
	fmt.Println(s)

}