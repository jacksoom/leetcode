package main

import (

	"fmt"
	//"sort"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.

func main() {
	s:="123DABCDABD"
	fmt.Println("result->",kmp(s,"DABD"))
}

func getNext(str string) []int {
	length:=len(str)
	next :=make([]int,length)
	j:=0
	k:=-1
	next[0]=-1
	for j<length-1 {
		//fmt.Println("j->",j,"k->",k)
		if k==-1 || str[j]==str[k] {
			//fmt.Println(string(str[j]))
			j++
			k++
			next[j]=k
		}else{
			k=next[k]
		}
	}
	return next
}

func kmp(str,sub string)int {
	if len(str)==0 || len(sub)==0{
		return -1
	}
	arr:=getNext(sub)
	j:=0
	var i int
	for i=0;i<len(str);i++ {
		fmt.Println(i,"--->",j)
		fmt.Println(string(str[i]),"--->",string(sub[j]))
		fmt.Println("___________________________________")
		for j>0 && str[i]!=sub[j]{
			j=arr[j-1]
		}
		if str[i]==sub[j] {
			j++
		}
		if len(sub)==j {
			return i-j+1
		}
	}
	return -1
}