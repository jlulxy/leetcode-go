package main

import "fmt"

/**
给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。

示例:

nums = [1, 2, 3]
target = 4

所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)

 */
 //分析为完全背包
 //思维错误因为要的是组合的结果所以【1,2】和【2，1】是一样的
/*
func combinationSum4(nums []int, target int) int {
	//dp[j] = max(dp[j],dp[j-i]+1)
	dp := make([]int, target+1)
	for i := 0; i < target+1; i++ {
		dp[i] = 0
	}
	dp[0]=1
	for i:=0 ;i<len(nums);i++{
		completePackage(dp,nums[i],1)
	}
	return dp[target]
}

func completePackage(dp[]int,cost,weight int ){
	for i := 0;i<len(dp) ;i++{
		if i - cost > 0{
			dp[i] += dp[i-cost]
		}
	}
}*/

/***
	上面的任务是完全背包想法错误因为没有考虑位置不同的结果陷入思维误区
	dp[i] = dp[i-cost1] + dp[i-cost2] + dp[i-cost3] ....
	从下往上dp推到，同时dp[0]代表空 空就什么都不拿随意dp【0】=1
*/
func combinationSum4(nums []int, target int) int {
	//dp[j] = max(dp[j],dp[j-i]+1)
	dp := make([]int, target+1)
	for i := 0; i < target+1; i++ {
		dp[i] = 0
	}
	dp[0]=1
	for i:=1; i<len(dp);i++{
		for j:=0;j<len(nums);j++{
			if i>=nums[j]{
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
func Max377(a,b int)int{
	if a>b {
		return a
	}
	return b
}

func main(){
	data :=[]int{1,2,3}
	target := 4
	fmt.Println(combinationSum4(data,target))
}
