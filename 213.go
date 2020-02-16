package main

import "fmt"


func rob1(nums []int) int {
	if len(nums)==0{
		return 0
	}
	//一开始basecase设为-10000 荒谬啦 偷东西还能赔钱哈哈
	dp_i__1 := 0
	dp_i_0  := 0
	for i:=0;i<len(nums);i++{
		temp := dp_i_0
		dp_i_0 = max213(dp_i_0,dp_i__1+nums[i])
		dp_i__1 = temp
	}
	return dp_i_0
}

func max213(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func rob(nums []int ) int {
	if len(nums) == 0{
		return 0
	}
	if len(nums) ==1 {
		return nums[0]
	}
	//互斥就是互相不考虑 可以计算不包含1或不包含尾的取最大
	return max213(rob1(nums[1:len(nums)]),rob1(nums[0:len(nums)-1]))
}

func main(){
	data := []int{1,2,3,2}
	fmt.Println(rob(data))
}