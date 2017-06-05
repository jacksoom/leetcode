package question

import (
	"strings"
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var maxLength int
	var strMap = make(map[string]int)
	length := strings.Count(s, "") - 1
	if length <= 0 {
		return 0
	}
	var i int
	var j int=0
	for i = 0; i < length; i++ {
		charAt := i
		//如果这个字符串在map中记录过,将其游标取代记录当前subString游标
		if _, ok := strMap[string(s[charAt])]; ok {
			j = max(j, strMap[string(s[charAt])]+1)
		}
		strMap[string(s[charAt])] = charAt
		maxLength = max(maxLength, charAt-j)
		fmt.Println(strMap)
	}
	return maxLength+1
}

func max(first int, args ...int) int {
	for _, v := range args {
		if first < v {
			first = v
		}
	}
	return first
}
func main(){
	fmt.Println(lengthOfLongestSubstring("abcacbdbb"))
}