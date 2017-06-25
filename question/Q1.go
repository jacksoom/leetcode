package main

import (
	"fmt"
	"reflect"
)

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
	var any interface{}
	s:=any.(Stringer)
	fmt.Println(reflect.TypeOf(s))
}
type Stringer interface {
	String() string
}
