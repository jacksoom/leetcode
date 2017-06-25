package main

import "fmt"

func main() {
	var i int8 = 5
	s := "abcdefgh"
	fmt.Println(s[i])
}

/*
	str-src
	str-pattern
*/
//addr为在src中坏字符串,距离模式串最大长度
func preBmBc(str string, failAddr int) [256]int {
	var result [256]int
	for i := 0; i < 255; i += 1 {
		result[i] = failAddr
	}
	//距离取代制
	for j := 0; j < failAddr-1; j += 1 {
		result[str[j]] = failAddr - j - 1
	}
	return result
}

func suffixes(str string, addr int) []int {
	var f int
	var suff = make([]int, addr)
	suff[addr-1] = addr
	g := addr-1
	for i := addr - 2; i >= 0; i -= 1 {
		if i>g && suff[i+addr-1-f]<i-g {
			suff[i]=suff[i+addr-1-f]
		}else{
			if i<g {
				g=i
			}
			f=i
			for g>=0&&str[g]==str[g+addr-1-f] {
				g--
			}
			suff[i]=f-g
		}
	}
	return suff
}
