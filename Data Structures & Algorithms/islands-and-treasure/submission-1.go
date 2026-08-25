func islandsAndTreasure(grid [][]int) {
	rows,cols:=len(grid),len(grid[0])
	queue:=[][2]int{}
	for i:=0;i<rows;i++{
		for j:=0;j<cols;j++{
			if grid[i][j]==0{
				queue=append(queue,[2]int{i,j})
			}
		}
	}
	dirs := [][2]int{
        {-1, 0},
        {1, 0},
        {0, -1},
        {0, 1},
    }
	for len(queue)!=0{
		r,c:=queue[0][0],queue[0][1]
		queue=queue[1:]
		for _,dir:=range dirs{
			if r+dir[0]<0 || r+dir[0]==rows || c+dir[1]<0 || c+dir[1]==cols{
				continue
			}
			if grid[r+dir[0]][c+dir[1]]!=2147483647{
				continue
			}
			grid[r+dir[0]][c+dir[1]]=grid[r][c]+1
			queue=append(queue,[2]int{r+dir[0],c+dir[1]})
		}
	}
    
}
