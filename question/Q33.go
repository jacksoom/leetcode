package main

import "fmt"

func search(arr []int,target int) int {
	length:=len(arr)
	if length<=0{
		return -1
	}
	left:=0
	right:=length-1
	for left<=right {
		mid:=(right+left)/2
		if arr[mid]==target {
			return mid
		}
		if target<arr[mid]{
			right=mid-1
		}else{
			left=mid+1
		}
	}
	return -1
}
func main()  {
	a:=[]int{1,3,5,7,9}
	fmt.Println(search(a,5))
}