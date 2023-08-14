package main

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Search(val int, curNode *TreeNode) *TreeNode {
	if curNode == nil {
		return nil
	} else if curNode.value == val {
		return curNode
	} else if val < curNode.value {
		return t.Search(val, curNode.left)
	} else {
		return t.Search(val, curNode.right)
	}
}
func main() {
}
