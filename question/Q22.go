package main

import "fmt"

func main() {
	str:=generateParenthesis(3)
	fmt.Println(str)
}
var (
	result []string
)
func generateParenthesis(n int) []string {
	generateOne("",n,n)
	return result
}
func generateOne(subStr string,left,right int){
	if left>right {
		return
	}
	if left>0{
		generateOne(subStr+"(",left-1,right)
	}
	if right>0{
		generateOne(subStr+")",left,right-1)
	}
	if left==0&&right==0{
		result=append(result,subStr)
		return
	}
}