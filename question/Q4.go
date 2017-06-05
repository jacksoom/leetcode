package main

import (
	"fmt"
	"math"
)
/*
	idea-1:判断num1.length+num2.length奇偶性,相当与找到,然后两个pointer在num1和num2循环查找，谁小谁++即可
	idea-2:
*/
func main() {
	a:=[]int{0,2,4}
	b:=[]int{1,3,5,6,8}
	fmt.Println(findMedianSortedArrays(a,b))
}
//
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var lengthNums1 = len(nums1)
	var lengthNums2 = len(nums2)
	var mid int
	mid = (lengthNums1+lengthNums2)/2
	//总数字个数为奇数时
	if (lengthNums1+lengthNums2)%2 != 0{
		return float64(findKth(nums1,0,nums2,0,mid+1))
	} else {
		return float64(findKth(nums1,0,nums2,0,mid)+findKth(nums1,0,nums2,0,mid+1))/2.0
	}
}
//类似二分查找思路
func findKth(array1 []int,aStart int,array2 []int,bStart int ,mid int) int {
	var lenArray1 = len(array1)
	var lenArray2 = len(array2)
	if aStart >= lenArray1 {
		return array2[bStart+mid-1]
	}
	if bStart >= lenArray2 {
		return array1[aStart+mid-1]
	}

	if mid == 1{
		return min(array1[aStart],array2[bStart])
	}
	var aKey,bKey int
	if aStart+mid/2-1 < lenArray1 {
		aKey=array1[aStart+mid/2-1]
	}else{
		aKey=math.MaxInt32
	}

	if bStart+mid/2 -1 <lenArray2 {
		bKey=array2[bStart+mid/2-1]
	}else{
		bKey=math.MaxInt32
	}

	if aKey<bKey{
		return findKth(array1,aStart+mid/2,array2,bStart,mid-mid/2)
	}else{
		return findKth(array1,aStart,array2,bStart+mid/2,mid-mid/2)
	}
}
func min(a int, b int) int {
	if a<b {
		return a
	}
	return b
}
