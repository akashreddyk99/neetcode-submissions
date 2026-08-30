func validTree(n int, edges [][]int) bool {
	adj:=make(map[int][]int)
	for _,v:=range edges{
		adj[v[0]]=append(adj[v[0]],v[1])
		adj[v[1]] = append(adj[v[1]], v[0])
	}
	visited:=make(map[int]bool)
	var dfs func(int,int)bool
	dfs=func(cur ,par int)bool{
		visited[cur]=true
		for _,child:=range adj[cur]{
			if child==par{
				continue
			}
			if visited[child]{
				return false
			}
			if !dfs(child,cur){
				return false
			}
		}
		return true
	}
	if !dfs(0,-1){
		return false
	}
	return len(visited)==n
    
}
