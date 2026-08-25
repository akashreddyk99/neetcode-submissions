func maxAreaOfIsland(grid [][]int) int {
	if len(grid)==0{
		return 0
	}
	ans:=0
	rows,cols:=len(grid),len(grid[0])
	visited:=make(map[[2]int]bool)
	var bfs func(int,int) int
	bfs= func(i,j int) int{
		queue:=[][2]int{}
		area:=0
		queue=append(queue,[2]int{i,j})
		visited[[2]int{i,j}]=true
		area+=1
		for len(queue)!=0{
			r,c:=queue[0][0],queue[0][1]
			queue=queue[1:]
			if r-1 >=0 && grid[r-1][c]==1 && !visited[[2]int{r-1,c}]{
				queue=append(queue,[2]int{r-1,c})
				visited[[2]int{r-1,c}]=true
				area+=1
			}
			if r+1 <rows && grid[r+1][c]==1 && !visited[[2]int{r+1,c}]{
				queue=append(queue,[2]int{r+1,c})
				visited[[2]int{r+1,c}]=true
				area+=1
			}
			if c-1 >=0 && grid[r][c-1]==1 && !visited[[2]int{r,c-1}]{
				queue=append(queue,[2]int{r,c-1})
				visited[[2]int{r,c-1}]=true
				area+=1
			}
			if c+1 <cols && grid[r][c+1]==1 && !visited[[2]int{r,c+1}]{
				queue=append(queue,[2]int{r,c+1})
				visited[[2]int{r,c+1}]=true
				area+=1
			}

		}
		return area
	}

	for i:=0;i<rows;i+=1{
		for j:=0;j<cols;j+=1{
			if grid[i][j]==1&&!visited[[2]int{i,j}]{
				area:=bfs(i,j)
				ans=max(ans,area)
			}
		}
	}
	return ans
}
