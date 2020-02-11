package main

func maxProfit(prices []int, fee int) int {
	if len(prices) == 0{
		return 0
	}

	dps:=make([]int,2)
	dps[0] = 0
	dps[1] = -prices[0]
	for i:=1;i<len(prices);i++{
		//卖的时候提交手续费，不能影响买的价格
		dps[0] = max121(dps[0],dps[1]+prices[i]-fee)
		//只能交易一次所以购买状态不依赖之前的收益
		dps[1] = max121(dps[1],dps[0]-prices[i])

	}
	return dps[0]
}

func maxProfitSpace (prices []int, fee int) int {
	if len(prices) == 0{
		return 0
	}
	cash ,hold := 0,intMin
	for i:=1;i<len(prices);i++ {
		//线更新cash和hold都可以 一个更新另一个一定不更新 数学太美啦
		cash = max121(cash,hold+prices[i]-fee)
		hold = max121(hold,cash-prices[i])
	}
	return cash
}

func max714(a,b int) int {
	if a>b {
		return a
	}
	return b
}