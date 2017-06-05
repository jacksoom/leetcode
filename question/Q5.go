package main

import (
	"fmt"
	"bytes"

	"github.com/ethereum/go-ethereum/common/math"
)

var maxLen int =0
func main() {
	extendPalidrome("122132442")
}

func longestPalindrome(s string) string {
	return ""
}
var T [50]int
func extendPalidrome(s string){
	length:=len(s)
	var buffer bytes.Buffer
	var i int
	for i=0;i<length;i++ {
		buffer.WriteString("#")
		buffer.WriteByte(s[i])
	}
	buffer.WriteString("#")
	str:=buffer.String()
	fmt.Println(str)
	length=len(str)
	var maxRigth int=1
	var maxId int=1
	var right int
	for i=1;i<length;i++{
		fmt.Println("loop start!")
		fmt.Println("maxright->>",maxRigth)
		if i>=maxRigth {
			fmt.Println("i->",i)
			fmt.Println("maxright->",maxRigth,"maxid->",maxId)
			fmt.Println("************************")
			T[i],right=findPalindromeInStr(str,i,length)
			fmt.Println("T[i]-->",T[i])
			fmt.Println("---------------------------------------------------------------------------------------")
			if right>=maxRigth{
				maxRigth=right-1
				maxId=i
			}
		}else{
			fmt.Println("i->",i)
			fmt.Println("maxright->",maxRigth,"maxid->",maxId)
			fmt.Println("++++++++++++++++++++++++")
			j:=2*maxId-i
			fmt.Println("j:",j)
			fmt.Println(j-T[j]," VS ",2*maxId-maxRigth)
			if j-T[j]>2*maxId-maxRigth{
				fmt.Println("else loop 1")
				T[i]=T[j]
			}else{
				fmt.Println("else loop 2")
				add := maxRigth-i
				T[i],right=find(str,i,maxRigth,length)
				T[i]+=add
				maxId=i
				maxRigth=right-1
			}
			fmt.Println("---------------------------------------------------------------------------------------")
		}
		math.MinInt32
	}
	fmt.Println(T)
}

func findPalindromeInStr(str string ,addr int,length int)(int,int)  {
	if addr==0 || addr == length-1 {
		return 0,addr
	}
	var left int
	var right int
	sum :=0
	left = addr-1
	right = addr+1
	for  {
		if left>=0&&right<=length-1 {
			fmt.Println(left,"-",right)
			if str[left]==str[right]{
				sum+=1
				right+=1
				left-=1
			}else{
				break
			}
		}else{
			break
		}

	}
	return sum,right
}
func find(str string,addr,rightStart,length int) (int,int){
	var sum int=0
	var left int
	var right int
	right=rightStart
	left=2*addr-rightStart
	for{
		if left>=0&&right<=length-1 {
			if str[left]==str[right]{
				sum+=1
				right+=1
				left-=1
			}else{
				break
			}
		}else{
			break
		}
	}
	return sum-1,right
}
func max(a int, b int) int {
	if a>b {
		return a
	}
	return b
}