func orangesRotting(grid [][]int) int {
	rows,cols:=len(grid),len(grid[0])
	queue:=[][3]int{}
	oc:=0
	for i:=0;i<rows;i++{
		for j:=0;j<cols;j++{
			if grid[i][j]==2{
				queue=append(queue,[3]int{i,j,0})
			}
			if grid[i][j]==1{
				oc+=1
			}
		}
	}

	dirs := [][2]int{
        {-1, 0},
        {1, 0},
        {0, -1},
        {0, 1},
    }
	time:=0

	for len(queue)!=0{
		r,c,level:=queue[0][0],queue[0][1],queue[0][2]
		queue=queue[1:]
		for _,dir:=range dirs{
			if r+dir[0]<0 || r+dir[0]==rows || c+dir[1] <0 || c+dir[1]==cols{
				continue
			}
			if grid[r+dir[0]][c+dir[1]]!=1{
				continue
			}
			grid[r+dir[0]][c+dir[1]]=2
			queue=append(queue,[3]int{r+dir[0],c+dir[1],level+1})
			oc-=1
			time=max(time,level+1)
		}
	}
	if oc!=0{
		return -1
	}
	return time
    
}
