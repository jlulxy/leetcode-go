package main

func coinChange(coins []int, amount int) int {
	//dp[i]当前总金额所需要的最小个数
	//dp[i]
	dp := make([]int,amount+1)
	//最开始没有物品时候消耗是0
	dp[0] = 0
	//因为要去最小值初始化一个相对的大值
	for i:=1;i<amount+1;i++{
		dp[i] = 10000
	}
	for i:=0;i<len(coins);i++{
		completePack(dp,coins[i],1)
	}
	return dp[amount]
}

func completePack (dp []int,cost int,w int){
	dpV := len(dp)
	for i:=0; i < dpV ; i++{
		if i >= cost{
			//通过转移方程dp[i] 和dp[i-cost] 保证amount被是被正常消耗完的
			dp[i] = minDp(dp[i],dp[i-cost]+w)
		}
	}
}
func minDp(a,b int) int{
	if a<b {
		return a
	}
	return b
}

func main(){
	data := []int{1,2,5}
	amout:=11
	coinChange(data,amout)
}