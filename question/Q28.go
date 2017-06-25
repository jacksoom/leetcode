package main

import "fmt"
//should use the KMP


func main() {
	fmt.Println(horspool("abcbcsdLibac-codecbcac","cbcac"))
}

func next(str string)[]int  {
	length:=len(str)
	next:=make([]int,length)
	k:=-1
	j:=0
	next[0]=-1
	for j<length-1{
		if(k==-1||str[j]==str[k]){
			k++
			j++
			next[j]=k
		}else{
			k=next[k]
		}
	}
	return next
}
func strStr(haystack,needle string) int{
	if (len(haystack)==0&&len(needle)==0)||len(needle)==0{
		return 0
	}
	if len(haystack)==0{
		return -1
	}
	i:=0
	j:=0
	sLen:=len(haystack)
	pLen:=len(needle)
	next:=next(needle)
	for i<sLen&&j<pLen {
		if j==-1||haystack[i]==needle[j]{
			i++
			j++
		}else{
			j=next[j]
		}
	}
	if j==pLen{
		return i-j
	}
	return -1
}

func horspool(src,pattern string) int {
	len1:=len(src)
	len2:=len(pattern)
	if len2==0||(len1==0&&len2==0){
		return 0
	}
	if len1==0{
		return -1
	}
	var i,j,k int
	var mispos int =0
	var mischar byte
	fmt.Println("start")
	for i=0;i<=len1-len2;{
		//fmt.Println("i->",i)
		for j=len2-1;j>=0;j-=1{
			if src[i+j]!=pattern[j] {
				//fmt.Println(src[i+j],"->",pattern[j])
				//fmt.Println("loop1")
				mischar=src[i+j]
				mispos=j
				break
			}
			//匹配成功
			if j==0 {
				return i
			}

		}
		//在pattern配字符串的左边开始寻找失陪字符串，并且并且开始跳跃寻找
		for k=mispos-1;k>=0;k-- {
			//fmt.Println("mispos->",mispos,"k->",k)
			if (pattern[k]==mischar){
				i+=(mispos-k)
				break
			}
			//如果没有失配字符,直接移动到src失配字符的前面
			if k==0{
				i+=(mispos+1)
				fmt.Println("i->",i)
			}

		}

	}
	return -1
}

