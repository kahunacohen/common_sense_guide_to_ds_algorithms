package main

type TreeNode struct {
	Value int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) searchRecur(val int, curNode *TreeNode) *TreeNode {
	if curNode == nil {
		return nil
	} else if curNode.Value == val {
		return curNode
	} else if val < curNode.Value {
		return t.searchRecur(val, curNode.left)
	} else {
		return t.searchRecur(val, curNode.right)
	}
}
func (t *TreeNode) Search(val int) *TreeNode {
	return t.searchRecur(val, t)
}
func main() {
}
