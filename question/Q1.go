package main
import "fmt"

//twoSum in leetcode question 1
func twoSum(nums []int, target int) []int {
	var result=make([]int,2)
	var map1 =make(map[int]int)
	for i,number := range nums{

		if _,ok:=map1[target-number];ok {
			result[1]=i
			result[0]=map1[target-number]
			return result
		}
		map1[number]=i
	}
	return result
}

func main(){
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b:=twoSum(a,17)
	fmt.Println(b[0],b[1])
}
