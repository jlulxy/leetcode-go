package main

import "fmt"

/**
	有交易限制 时间o(n)空间o(n)
	时间 100%
	空间 16.00%
 */
func maxProfit309(prices []int) int {
	//dp[i][0] = max(dp[i-1][0],dp[i-1][1]+price[i])
	//dp[i][1] = max(dp[i-1][1],dp[i-2][0] -price[i])
	if len(prices) <= 1{
		return 0
	}
	dp:= make([][]int,len(prices))
	for i:=0;i<len(prices);i++{
		dp[i] = make([]int,2)
	}
	dp[0][1] = -prices[0]
	//要注意第一手买的时候不需要有隔天限制要判断下base数据在第一天买便宜还是在第二天买便宜
	dp[1][1] = max309(-prices[1],-prices[0])
	for i := 1;i<len(prices);i++{
		dp[i][0] = max309(dp[i-1][0],dp[i-1][1]+prices[i])
		if i-2 >= 0{
			dp[i][1] = max309(dp[i-1][1],dp[i-2][0] -prices[i])
		}
	}
	return dp[len(prices)-1][0]
}

/**
	有交易限制 时间o(n)空间o(1)
	时间 100%
	空间 80.00%
 */
func maxProfitSpace(prices []int) int {
	//dp[i][0] = max(dp[i-1][0],dp[i-1][1]+price[i])
	//dp[i][1] = max(dp[i-1][1],dp[i-2][0] -price[i])
	if len(prices) ==0 {
		return 0
	}
	dp_i_0 := 0
	dp_i_1 := -10000
	dp_pre := 0
	for i:=0;i<len(prices);i++{
		temp := dp_i_0
		dp_i_0 = max309(dp_i_0,dp_i_1+prices[i])
		dp_i_1 = max309(dp_i_1,dp_pre-prices[i])
		dp_pre = temp
	}
	return dp_i_0
}


func max309(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func main(){
	data := []int{1,2,4}
	fmt.Println(maxProfitSpace(data))
}