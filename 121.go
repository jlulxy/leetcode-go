package main

import "fmt"

/**
   未优化 标准转移方程 要求只能交易一次
 */
func maxProfitDp(prices []int) int {
	//dp 定义
	//dp[i][0] = max(dp[i-1][0],dp[i-1][1] + prices[i])
	//dp[i][1] = max(dp[i-1][1],dp[i-1][0] - prices[i])
	if len(prices) == 0{
		return 0
	}
	dps := make([][]int,len(prices)+1)
	for i:=0;i<=len(prices);i++{
		dps[i] = make([]int,2)
		for j:=0;j<2;j++{
			if i == 0{
				dps[i][j] = 0
			}
		}
	}
	dps[0][0] = 0;
	dps[0][1] = -prices[0];
	for i:=1;i<len(prices);i++{
		dps[i][0] = max121(dps[i-1][0],dps[i-1][1]+prices[i])
		//只能交易一次所以购买状态不依赖之前的收益
		dps[i][1] = max121(dps[i-1][1],-prices[i])

	}

	return dps[len(prices)-1][0]
}

/***
	记录当前天为截至的最小价格，用当天价格去见如果差价大更新最大价格
	time 97%
	space 100%
 */
 func maxProfit(prices []int) int{
 	if len(prices) == 0{
 		return 0
	}
 	maxProfit,minData:=0,prices[0]

 	for i:=0;i<len(prices);i++{
 		if minData >prices[i] {
 			minData = prices[i]
		}
 		if prices[i] - minData >maxProfit{
 			maxProfit = prices[i] - minData
		}
	}
	return maxProfit
 }


func max121(a,b int) int {
	if a>b {
		return a
	}
	return b
}


func main(){
	data := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(data))
}