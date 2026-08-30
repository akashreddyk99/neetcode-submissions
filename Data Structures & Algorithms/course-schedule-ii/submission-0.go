func findOrder(numCourses int, prerequisites [][]int) []int {
	preMap:=make(map[int][]int)
	for _,v:=range prerequisites{
		preMap[v[0]]=append(preMap[v[0]],v[1])
	}
	ans:=[]int{}
	cycle:=make(map[int]bool)
	visited:=make(map[int]bool)
	var dfs func(int) bool
    dfs=func(crs int) bool{
		if cycle[crs]{
			return false
		}
		if visited[crs]{
			return true
		}
		cycle[crs]=true
		for _,pre:= range preMap[crs]{
			if !dfs(pre){
				return false
			}
		}
		delete(cycle,crs)
		visited[crs]=true
		ans=append(ans,crs)
		return true
	}
	for i:=0;i<numCourses;i++{
		if !dfs(i){
			return []int{}
		}
	}
	return ans
    
}
