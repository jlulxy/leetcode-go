package main

import "fmt"

func uniquePaths(m int, n int) int {
	//动态转移方程 F(m,n) = F(m,n-1)+F(m-1,n)\
	//F(m,n) =1 (m=0||n=0)
	fn := make([][]int,1)
	fn[0]= make([]int,n)
	pre := -1
	for i :=0;i<m;i++{
		for j:=0;j<n;j++{
			if i==0|| j==0{
				fn[0][j] = 1
			}else {
				fn[0][j] +=  pre
			}
			pre = fn[0][j]
		}
	}
	return fn[0][n-1]
}
func main()  {
	fmt.Println(uniquePaths(7,3))
	fmt.Println(uniquePaths(3,2))
}
