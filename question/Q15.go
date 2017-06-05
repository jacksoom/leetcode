package main

import "fmt"

func main() {
	value := []int{-1,0,1,2,-1,-4}

	fmt.Println(threeSum(value))
}

func threeSum(nums []int) [][]int {
	if len(nums)<3{
		return [][]int{}
	}
	var sli [][]int
	var map1 =make(map[int]int)
	quickSort(nums,0,len(nums)-1)
	for i,v := range nums{
		if v>0 {
			break
		}
		if _,ok:=map1[v];ok {
			continue
		}
		map1[v]=i
		res:=findEr(nums,i+1,len(nums)-1,v)
		if len(res)==0{
			continue
		}
		for k,val := range res {
			sli=append(sli,[]int{v,k,val})
		}
	}
	return sli
}

func quickSort(arr []int, start int, end int) {

	if len(arr)==0{
		return
	}

	var key  int = arr[start]
	var low  int = start
	var high int = end

	for {
		for low < high {
			if arr[high] < key {
				arr[low] = arr[high]
				break
			}
			high--
		}
		for low < high {
			if arr[low] > key {
				arr[high] = arr[low]
				break
			}
			low++
		}
		if low >= high {
			arr[low] = key
			break
		}
	}
	if low-1 > start {
		quickSort(arr, start, low-1)
	}
	if high+1 < end {
		quickSort(arr, high+1, end)
	}
}

func findEr(arr []int,left,right,target int)map[int]int{
	var map1 = make(map[int]int)
	for left <right {
		if arr[left]+arr[right]+target ==0 {
			map1[arr[left]]=arr[right]
			for left<right&&arr[left]==arr[left+1]{
				left++
			}
			for left<right && arr[right]==arr[right-1]{
				right--
			}
			left++
			right--
		}else if arr[left]+arr[right]+target<0{
			left++
		}else{
			right--
		}
	}
	return map1
}

func tSum(arr []int,left,right,target int)map[int]int {
	var map1 = make(map[int]int)
	var result = make(map[int]int)
	var i int
	for i = left; i < right+1; i++ {
		if _, ok := map1[target - arr[i]]; ok {
			result[target - arr[i]] = arr[i]
		}
		map1[arr[i]] = i
	}
	return result
}

