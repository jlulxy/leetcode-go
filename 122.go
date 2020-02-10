package main

import "fmt"

/*
	股票问题基本公式做空间压缩
	time ：12.88%
	space: 59.61%
*/

func maxProfitDpAll(prices []int) int {
	//dp 定义
	//dp[i][0] = max(dp[i-1][0],dp[i-1][1] + prices[i])
	//dp[i][1] = max(dp[i-1][1],dp[i-1][0] - prices[i])
	if len(prices) == 0{
		return 0
	}

	dps:=make([]int,2)
	dps[0] = 0
	dps[1] = -prices[0]
	for i:=1;i<len(prices);i++{
		dps[0] = max121(dps[0],dps[1]+prices[i])
		//只能交易一次所以购买状态不依赖之前的收益
		dps[1] = max121(dps[1],dps[0]-prices[i])

	}
	return dps[0]
}

/*
	所有大于前一天的数都加起来
	time 100%
	space 59.81%
*/
func maxProfitWF(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	data := 0
	for i:=1;i<len(prices);i++{
		if prices[i] > prices[i-1]{
			data += prices[i] -prices[i-1]
		}
	}
	return data
}

/**
	峰谷相减法
 	time 96.71%
	space 100%
 */
 func maxProfileFG(prices[]int) int {
 	if len(prices) == 0{
 		return 0
	}
	maxprofit , valley,peek, i := 0,prices[0],prices[0] ,0
	for i<len(prices)-1{
		for i< len(prices)-1 && prices[i] >=prices[i+1]{
			i++
		}
		valley = prices[i]
		for i< len(prices)-1 && prices[i] <= prices[i+1]{
			i++
		}
		peek = prices[i]
		maxprofit += peek -valley
	}
	return  maxprofit
 }

func main(){
	data := []int{7,1,5,3,6,4}
	fmt.Println(maxProfitFG(data))
}