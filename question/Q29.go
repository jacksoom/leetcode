package main

import (
	"math"
	"fmt"
)

func divide(dividend int, divisor int) int {
	if divisor==0 || (dividend==math.MinInt32&&divisor==-1) {
		return math.MaxInt32
	}
	var sign bool=(divisor>0 && dividend>0)||(divisor<0&&dividend<0)
	var A uint
	var B uint
	if dividend<0{
		A=uint(-dividend)
	}else{
		A=uint(dividend)
	}

	if divisor<0{
		B=uint(-divisor)
	}else{
		B=uint(divisor)
	}

	var result int=0
	var i int
	for i=31;i>=0;i--{
		if (A>>uint(i)) >= B{
			result=result<<1+1
			A-=(B<<uint(i))
		}else{
			result=result<<1
		}
	}
	if !sign {
		result=-result
	}
	return result
}

func main()  {
	fmt.Println(divide(9,2))
}