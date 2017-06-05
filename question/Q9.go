package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(12131321))
}
//中介反转
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	p := x
	q := 0
	for p >=10 {
		q =q*10+p%10
		p/=10
	}
	return q==x/10 && p==x%10
}