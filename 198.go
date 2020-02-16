package main

import "fmt"

func rob(nums []int) int {
	if len(nums)==0{
		return 0
	}
	//一开始basecase设为-10000 荒谬啦 偷东西还能赔钱哈哈
	dp_i__1 := 0
	dp_i_0  := 0
	for i:=0;i<len(nums);i++{
		temp := dp_i_0
		dp_i_0 = max198(dp_i_0,dp_i__1+nums[i])
		dp_i__1 = temp
	}
	return dp_i_0
}

func max198(a,b int) int {
	if a>b {
		return a
	}
	return b
}
func main(){
	data:=[]int{2,1,5}
	fmt.Println(rob(data))
}