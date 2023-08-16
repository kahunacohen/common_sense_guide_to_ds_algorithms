package main

import "fmt"

type TreeNode struct {
	Value int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) searchRecur(val int, curNode *TreeNode) *TreeNode {
	fmt.Printf("Searching: %d\n", curNode.Value)
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
func (t *TreeNode) insertRecur(val int, prevNode, curNode *TreeNode) {
	if curNode == nil {
		newNode := TreeNode{Value: val}
		if val < prevNode.Value {
			prevNode.left = &newNode
		} else {
			prevNode.right = &newNode
		}
		return
	}
	if val < curNode.Value {
		t.insertRecur(val, curNode, curNode.left)
	}
	if val > curNode.Value {
		t.insertRecur(val, curNode, curNode.right)
	}
}
func (t *TreeNode) Insert(val int) {
	t.insertRecur(val, nil, t)
}
func main() {
}
