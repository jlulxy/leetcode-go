package main

import "fmt"

//返回costdp 01的个数
func CheckStr(strs []string) [][2]int{
	costString := make([][2]int,len(strs))
	for key ,str := range strs{
		cout:=0
		for _,char := range str{
			if string(char) == "0"{
				cout++
			}
		}
		costString[key][0] = cout
		costString[key][1] = len(str) - cout
	}
	return costString
}

func findMaxForm(strs []string, m int, n int) int {
	//costArray := CheckStr(strs)
	dp := make([][]int,m+1)
	for i:=0;i<m+1;i++{
		dp[i] = make([]int,n+1)
	}
	one, zero := 0, 0
	for i:=0;i<len(strs);i++{
		one, zero = 0, 0
		for _, char := range strs[i] {
			if char == '0' {
				zero++
			} else {
				one++
			}
		}
		for j:=m;j>=0;j--{
			for k:=n;k>=0;k--{
				if j- zero >=0 && k-one>=0{
					dp[j][k] = max474(dp[j][k],dp[j-zero][k-one]+1)
				}

			}
		}
	}
	return dp[m][n]

}

func max474(a,b int)int{
	if a>b{
		return a
	}
	return b
}


func main ()  {
	data:=[]string{"10","0","1"}
	fmt.Println(findMaxForm(data,1,1))
}
