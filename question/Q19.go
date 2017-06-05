package main

import (
	"fmt"

)

func main() {
	a1:=ListNode{1,nil}

	//result:=&a1
	//1 2 3 4 5
	res:=removeNthFromEnd(&a1,1)
	for res!=nil {
		fmt.Println(res.Val)
		res=res.Next
	}
}
type ListNode struct{
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l1:=head
	l2:=head
	for l1 != nil{
		l1=l1.Next
		if n <0{
			l2=l2.Next
		}
		n--
	}
	if n==0 {
		head=head.Next
	}else{
		l2.Next=l2.Next.Next
	}
	return head
}