func pacificAtlantic(heights [][]int) [][]int {
	rows,cols:=len(heights),len(heights[0])
	pas:=make(map[[2]int]bool)
	atl:=make(map[[2]int]bool)
	var dfs func(int,int,int, map[[2]int]bool)

	dfs= func(i,j,minheight int,visit map[[2]int]bool){
		if i<0 ||j<0 || i==rows||j==cols || visit[[2]int{i,j}] || heights[i][j]<minheight{
			return
		}
		visit[[2]int{i,j}]=true
		dfs(i+1,j,heights[i][j],visit)
		dfs(i-1,j,heights[i][j],visit)
		dfs(i,j+1,heights[i][j],visit)
		dfs(i,j-1,heights[i][j],visit)
	}
	for c:=0;c<cols;c++{
		dfs(0,c,heights[0][c],pas)
		dfs(rows-1,c,heights[rows-1][c],atl)
	}
	for r:=0;r<rows;r++{
		dfs(r,0,heights[r][0],pas)
		dfs(r,cols-1,heights[r][cols-1],atl)
	}
	ans:=[][]int{}
	for k,_:=range pas{
		if atl[k]{
			ans=append(ans,[]int{k[0],k[1]})
		}
	}
	return ans

	
}
