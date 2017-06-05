package main

import "fmt"

func main() {
	a := [3]int{1,2,1}
	fmt.Println(maxArea(a[:]))
}

func maxArea(height []int) int {
	length:= len(height)
	left := 0
	right := length-1
	max :=0
	for left<right {
		max=maxInt(max,minInt(height[left],height[right])*(right-left))
		if height[left]<height[right] {
			left++
		}else{
			right--
		}
	}
	return max
}

func maxInt(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func minInt(a,b int) int {
	if a>b {
		return b
	}
	return a
}