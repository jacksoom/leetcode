package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lnode1 := l1
	lnode2 := l2
	sentinel := newNode(0)
	node := sentinel
	sum := 0
	for {
		if lnode1 == nil && lnode2 == nil {
			break
		}
		sum /= 10
		if lnode1 != nil {
			sum += lnode1.Val
			lnode1 = lnode1.Next
		}
		if lnode2 != nil {
			sum += lnode2.Val
			lnode2 = lnode2.Next
		}
		node.Next = newNode(sum % 10)
		node = node.Next
	}
	if sum/10 == 1 {
		node.Next = newNode(1)
	}
	return sentinel.Next
}

func newNode(value int) *ListNode {
	return &ListNode{Val: value, Next: nil}
}

func printList(node *ListNode) {
	lnode := node
	for {
		if lnode == nil {
			break
		}
		fmt.Print(lnode.Val)
		lnode = lnode.Next
	}
	fmt.Println("")
}
func main() {
	node1 := newNode(2)
	node2 := newNode(4)
	node3 := newNode(9)

	node4 := newNode(5)
	node5 := newNode(6)
	node6 := newNode(4)
	node7 := newNode(9)

	node1.Next = node2
	node2.Next = node3

	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	printList(node1)

	printList(node4)
	result := addTwoNumbers(node1, node4)

	printList(result)
}
