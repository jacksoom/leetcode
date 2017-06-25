package main
func removeElement(nums []int, val int) int {
	length:=len(nums)
	i:=0
	j:=length-1
	for i<length {
		if nums[i]==val{
			t:=nums[j]
			nums[j]=nums[i]
			nums[i]=t
			length--
			j--
		}else{
			i++
		}
	}

	return length
}
