package main


func canPartition(nums []int) bool {
	//取和判断终止条件
	sum :=0
	for k := range nums{
		sum += nums[k]
	}

	//奇数直接返回
	if  (sum & 1) == 1{
		return false
	}
	target :=sum/2
	//dp[i][j] 在数组[0-i]恰巧可以等于 选取等于j的结果
	dp := make([][]bool,len(nums))
	for i:=0;i< len(nums);i++{
		dp[i] = make([]bool,target+1)
		for j:=0;j < target+1 ;j++{
			dp[i][j] = false
		}
	}
	//开启转移
	if (nums[0]<=target){
		dp[0][nums[0]] = true
	}
	for i:=1; i< len(nums);i++{
		for j:=0;j<=target;j++{
			dp[i][j] = dp[i-1][j]
			if nums[i] == j {
				dp[i][j] = true
				continue;
			}  else if(nums[i]<j){
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][target]
}

//优化空间
func canPartitionV2 (nums []int) bool {
	//取和判断终止条件
	sum :=0
	for k := range nums{
		sum += nums[k]
	}

	//奇数直接返回
	if  (sum & 1) == 1{
		return false
	}
	target :=sum/2
	//[j] 在数组[0-i]恰巧可以等于 选取等于j的结果
	dp := make([]bool,target +1)
	for i:=0;i< target+1;i++{
		dp[i] = false
	}
	dp[0] = true
	//i循环可以看作不断扩大范围，dp[j]表示在当前id范围下符合条件的dp值，
	// 根据dp可以发现j从大到下计算保证计算的因子都是上次循环计算过的

	for i:= 0; i< len(nums);i++{
		for j:=target;j>=nums[i];j--{
			if (dp[target]){
				return true
			}
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}
