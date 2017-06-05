package main

import (
	"bytes"
	"fmt"
)

func intToRoman(num int) string {
	value := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	str := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	var buffer bytes.Buffer
	for k, v := range value {
		for num>=v {
			num-=v
			buffer.WriteString(str[k])
		}
	}
	return buffer.String()
}

func main() {
	fmt.Println(intToRoman(1234))
}