package main

import (

	"fmt"
	"bytes"
	"strconv"
)

func main() {
	myAtoi("1 2 3")
}
func myAtoi(str string) int {
	var buffer bytes.Buffer
	var i int
	for i=0;i<len(str);i++ {
		if str[i]!=32 {
			buffer.WriteByte(str[i])
		}
	}
	str1:=buffer.String()
	strconv.ParseInt(str1, 10, )
	return 0
}
