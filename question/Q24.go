package main


type ListNode struct {
     Val int
     Next *ListNode
 }

func swapPairs(head *ListNode) *ListNode {
	var p,q,m *ListNode
	if head==nil||head.Next==nil{
		return head
	}
	//swap head and head.next
	p=head.Next
	q=p.Next
	head.Next=q
	p.Next=head
	head=p
	p=head.Next
	//whethre there are two nodes behind
	for haveTwo(p){
		q=p.Next
		m=q.Next
		q.Next=m.Next
		m.Next=q
		p.Next=m
		p=q
	}
	return head
}
func haveTwo(node *ListNode) bool {
	if node.Next!=nil&&node.Next.Next!=nil{
		return true
	}
	return false
}
