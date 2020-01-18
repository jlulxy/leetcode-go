package main

import "fmt"

//思维就是错的！！！！
func lengthOfLISWrong(nums []int) int {
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		tempLen := 1
		pre := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > pre {
				tempLen++
				pre = nums[j]
			}

		}
		if tempLen > maxLen {
			maxLen = tempLen
		}
	}
	return maxLen
}

//动态规划：if nums[i] >nums[j] 0<j<i
//dp[i]=max(dp[i],dp[j]+1)
func lengthOfLIS(nums []int) int {
	dp := make([]int,len(nums))
	maxLen := 0
	if len(nums) == 0{
		return maxLen
	}
	//初始化为1
	for i := 0 ;i <len(dp);i++{
		dp[i] = 1
	}
	for i:= 1;i<len(nums);i++ {
		for j:=0;j<i;j++{
			if nums[i] >nums[j] {
				dp[i] = max(dp[i],dp[j]+1)
			}
		}
		if dp[i]> maxLen{
			maxLen = dp[i]
		}
	}
	return maxLen
}
func max(a,b int) int {
	if a>b{
		return a
	}
	return b
}


func main() {
	input := []int{10,9,2,5,3,4}
	fmt.Println(lengthOfLIS(input))
}
