func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap:=make(map[int][]int)
	for _,v :=range prerequisites{
		preMap[v[0]]=append(preMap[v[0]],v[1])
	}

	visitMap:=make(map[int]bool)
	var dfs func(int)bool
	dfs=func(crs int) bool{
		if visitMap[crs]{
			return false
		}
		if len(preMap[crs])==0{
			return true
		}
		visitMap[crs]=true
		for _,pre:=range preMap[crs]{
			if !dfs(pre){
				return false
			}
		}
		delete(visitMap,crs)
		preMap[crs]=[]int{}

		return true
	}

	for crs:=0;crs<numCourses;crs++{
		if !dfs(crs){
			return false
		}
	}
	return true
}
