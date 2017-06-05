package main

import (
	"fmt"

)

func main() {
	value := []int{-1,0,1,2,-1,-4,9}
	quickSort(value,0,len(value)-1)
	fmt.Println(value)
}
func quickSort(sortArray []int, left, right int) {
	if left < right {
		pos := partition(sortArray, left, right)
		quickSort(sortArray, left, pos-1)
		quickSort(sortArray, pos+1, right)
	}
}

func partition(arr []int,left,right int)int{
	if len(arr)==0{
		return -1
	}
	pivot := arr[left]
	for left<right {
		for left<right && arr[right]>=arr[left]{
			right--
		}
		if left<right {
			arr[left]=arr[right]
			left++
		}
		for left<right && arr[left]<=pivot{
			left++
		}
		if left<right{
			arr[right]=arr[left]
			right--
		}
	}
	arr[left]=pivot
	return left
}

// Swap the position of a and b
func swap(arr []int,a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}