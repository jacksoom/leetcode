package main

import "fmt"

func main() {
	a:=ListNode{1,nil}
	b:=ListNode{2,nil}
	c:=mergeTwoLists(&a,&b)
	for c!=nil{
		fmt.Println(c.Val)
		c=c.Next
	}
}


type ListNode struct {
      Val int
      Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = new(ListNode)
	var result *ListNode = head
	for l1!=nil&&l2!=nil {
		if l1.Val>l2.Val{
			result.Next=l2
			l2=l2.Next
		}else{
			result.Next=l1
			l1=l1.Next
		}
		result=result.Next
	}
	if l1==nil{
		result.Next=l2
	}else{
		result.Next=l1
	}
	return head.Next
}