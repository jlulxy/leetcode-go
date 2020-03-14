package main

import "fmt"

// 思路：从00开始 上下左右有1的时候就遍历 寻找为1的联通预支，找到后就标记已经处理不在进行重复计算。
func numIslands(grid [][]string) int {
	if len(grid) == 0{
		return 0
	}
	w := len(grid)
	h := len(grid[0])
	cout:=0
	var dealRecord = make([][]bool,w)
	for k :=range dealRecord{
		dealRecord[k] = make([]bool,h)
	}
	for i:=0;i<w;i++{
		for j:=0;j<h;j++{
			if grid[i][j] == "1" && !dealRecord[i][j] {
				cout++
				numIslandsre(i,j,grid,dealRecord)
			}
		}
	}
	return cout
}

// numIslandsre 遍历上下左右对联通的打上标记
func numIslandsre(i,j int,grid [][]string,record [][]bool) {
	//剪枝
	if record[i][j]{
		return
	}
	record[i][j] = true
	if grid[i][j] == "0"{
		return
	}
	if j+1<len(grid[0]) && !record[i][j+1]{
		numIslandsre(i,j+1,grid,record)
	}
	if i+1<len(grid) && !record[i+1][j]{
		numIslandsre(i+1,j,grid,record)
	}
	if j-1>=0 && !record[i][j-1]{
		numIslandsre(i,j-1,grid,record)
	}
	if i-1>=0 && !record[i-1][j]{
		numIslandsre(i-1,j,grid,record)
	}
	return
}

func main(){
	data := [][]string{{"1","1","1"},{"0","1","0"},{"1","1","1"}}
	re:=numIslands(data)
	fmt.Println(re)
}
