func numIslands(grid [][]byte) int {
	if len(grid)==0{
		return 0
	}
	rows,col:=len(grid),len(grid[0])
	visited := make(map[[2]int]bool)
	ans:=0
	var bfs func(int,int)
	bfs=func(i,j int){
		queue:=[][2]int{}
		visited[[2]int{i,j}]=true
		queue=append(queue,[2]int{i,j})
		for len(queue)!=0{
			r,c:=queue[0][0],queue[0][1]
			queue=queue[1:len(queue)]
			if r-1>=0 && grid[r-1][c]=='1' && !visited[[2]int{r-1,c}]{
				queue=append(queue,[2]int{r-1,c})
				visited[[2]int{r-1,c}]=true

			}
			if r+1<rows && grid[r+1][c]=='1'&&!visited[[2]int{r+1,c}]{
				queue=append(queue,[2]int{r+1,c})
				visited[[2]int{r+1,c}]=true

			}
			if c-1>=0 && grid[r][c-1]=='1'&&!visited[[2]int{r,c-1}]{
				queue=append(queue,[2]int{r,c-1})
				visited[[2]int{r,c-1}]=true

			}
			if c+1<col && grid[r][c+1]=='1'&&!visited[[2]int{r,c+1}]{
				queue=append(queue,[2]int{r,c+1})
				visited[[2]int{r,c+1}]=true

			}
		}
	}
	for i:=0;i<rows;i++{
		for j:=0;j<col;j++{
			if grid[i][j]=='1'&&!visited[[2]int{i,j}]{
				bfs(i,j)
				ans+=1
			}
		}
	}
	return ans
}
