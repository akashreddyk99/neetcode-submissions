func solve(board [][]byte) {
	rows,cols:=len(board),len(board[0])
	var dfs func(int,int,map[[2]int]bool) bool
	dfs= func(i,j int,visited map[[2]int]bool) bool{
		if i<0||i==rows||j<0||j==cols{
			return false
		}else if board[i][j]=='O'{
			if visited[[2]int{i,j}]{
			return true
			}
			visited[[2]int{i,j}]=true
			a:=dfs(i+1,j,visited)
			b:=dfs(i-1,j,visited)
			c:=dfs(i,j+1,visited)
			d:=dfs(i,j-1,visited)
			return a&&b&&c&&d
		}
		return true
	}
	for i:=0;i<rows;i++{
		for j:=0;j<cols;j++{
			if board[i][j]=='O'{
				visited:=make(map[[2]int]bool)
				if dfs(i,j,visited){
					for k,_:=range visited{
						board[k[0]][k[1]]='X'
					}
				}
			}
		}
	}
}
