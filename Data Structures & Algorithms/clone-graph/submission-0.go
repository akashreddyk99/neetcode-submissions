/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node==nil{
		return node
	}
	oldtoNew:=make(map[*Node]*Node)
	var clone func(*Node) *Node
	clone= func(node *Node) *Node{
		if _,ok:=oldtoNew[node];ok{
			return oldtoNew[node]
		}
		copy:=&Node{
			Val:node.Val,
			Neighbors: []*Node{},
		}
		oldtoNew[node]=copy
		for _,child:=range node.Neighbors{
			copy.Neighbors=append(copy.Neighbors,clone(child))
		}
		return copy
	}
	return clone(node)
}