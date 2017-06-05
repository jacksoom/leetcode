package main

import (
	"bytes"
	"fmt"
	"strconv"
	"math"
)

func main() {
	fmt.Println(byte())
}

func reverse(x int) int {
	if x==0{
		return 0
	}
	if x<0 {
		num, err :=strconv.ParseInt(reverseStr(strconv.Itoa(-x)), 10, 64)
		if err!=nil{
			return 0
		}
		if num > math.MaxInt32{
			return 0
		}

		return -int(num)
	}else{
		num, err :=strconv.ParseInt(reverseStr(strconv.Itoa(x)), 10, 64)
		if err!=nil{
			return 0
		}
		if num > math.MaxInt32{
			return 0
		}

		return int(num)
	}

}
//int to string
func reverseStr(s string)string{
	var buffer bytes.Buffer
	var i int
	addr :=findNotZero(s)
	for i=addr;i>=0;i--{
		buffer.WriteByte(s[i])
	}
	return buffer.String()
}
func findNotZero(s string) int {
	var i int
	for i=len(s)-1;i>=0;i++ {
		if s[i]!=0 {
			return i
		}
	}
	return 0
}