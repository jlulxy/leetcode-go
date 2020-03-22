package main

import (
	"fmt"
	"container/list"
)


type data struct {
	X int
	Y int
}
// maxDistance 暴力求解 超出时间,失去了图的状态
func maxDistanceBao(grid [][]int) int {
	lands := []data{}
	seas  := []data{}
	for i,item := range grid{
		for j ,v := range item{
			tmp := data{
				X:i,
				Y:j,
			}
			if v == 1  {
				lands = append(lands,tmp)
			}else{
				seas  = append(seas,tmp)
			}
		}
	}
	if len(lands) ==0 ||len(seas) == 0{
		return -1
	}
	// 记录每个海域和
	maxDistance := -1
	for _,sea :=range seas{
		tempMinDistance := 2000
		for _,land := range lands{
			distance := abs(sea.X-land.X) + abs(sea.Y-land.Y)
			if distance < tempMinDistance{
				tempMinDistance = distance
			}
		}
		if tempMinDistance > maxDistance {
			maxDistance = tempMinDistance
		}
	}
	return maxDistance
}

type viaRecord struct {
	X  int
	Y  int
	Sep int
}

var (
	dx = [4]int{-1,0,1,0}
	dy = [4]int{0,1,0,-1}
)


func maxDistanceWFS(grid [][]int) int {
	ans := -1
	if len(grid)== 0 {
		return -1
	}
	n := len(grid)
	m := len(grid[0])
	for i,item := range grid {
		for j,v := range item{
			if v == 0 {
				fmt.Println("start",i,j)
				ans = max1162(ans,findNearestLandWFS(grid,i,j,n,m))
			}
		}
	}
	return ans
}

// 宽度优先遍历 每次重建遍历记录表，在同一step下找到即可返回 O(4)
func findNearestLandWFS(grid [][]int,x,y int,n,m int) int {

	var queue = list.New()
	queue.PushBack(viaRecord{x,y,0})
	var vis [100][100]int
	vis[x][y] = 1
	for queue.Len()!= 0{
		f := queue.Front()
		queue.Remove(queue.Front())
		if record,ok := f.Value.(viaRecord);ok{
			for i:=0;i<4;i++{
				nx := record.X + dx[i]
				ny := record.Y + dy[i]
				sep := record.Sep +1
				if !(nx >= 0 && ny >=0 && nx <= n-1 &&ny<= m-1){
					continue
				}
				if vis[nx][ny] != 1 {
					queue.PushBack(viaRecord{nx,ny,sep})
					vis[nx][ny] = 1;
					if grid[nx][ny] == 1{
						return sep
					}
				}
			}
		}
	}
	return -1;
}


type viaRecordfill struct {
	X int
	Y int
}

// 填海造陆地 需要多少步填满海就有多少最大距离
func maxDistance(grid [][]int) int {
	if len(grid)== 0 {
		return -1
	}
	n := len(grid)
	m := len(grid[0])
	que :=list.New()
	for i,item :=range grid{
		for j,v := range item{
			if v == 1{
				que.PushBack(viaRecordfill{i,j})
			}
		}
	}
	level := -1
	for que.Len() != 0{
		key := que.Len()
		fill := false
		for i:=0;i<key;i++{
			f := que.Front()
			que.Remove(que.Front())
			if record,ok := f.Value.(viaRecordfill);ok{
				for j :=0; j<4; j++{
					nx := record.X + dx[j]
					ny := record.Y + dy[j]
					if nx >= 0 && ny >=0 && nx <= n-1 &&ny<= m-1 && grid[nx][ny] == 0{
						que.PushBack(viaRecordfill{nx,ny})
						grid[nx][ny] = 1
						fill = true
					}
				}

			}
		}
		if fill{
			level++
		}
	}
	return level
}

func abs(n int) int {
	if n < 0{
		return -n
	}
	return n
}

func max1162(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func main (){
	data:=[][]int{{1,0,1},{0,0,0},{1,0,1}}
	fmt.Println(maxDistance(data))
	fmt.Println("!!!!!")
	//fmt.Println(findNearestLand(data,2,2,3,3))
}