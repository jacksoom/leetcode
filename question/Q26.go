package main

import "fmt"

func main() {
	a:=[]int{1,1,2}
	result:=removeDuplicates(a)
	fmt.Println(result)
	fmt.Println(a)
}

func removeDuplicates(nums []int) int {
	length:=len(nums)
	if length==0{
		return 0
	}
	var i int
	j:=0
	for i=0;i<len(nums);i++{
		if nums[i]!=nums[j]{
			nums[j]=nums[i]
			j++
		}
	}
	j++
	return j
}