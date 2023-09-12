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

func (t *TreeNode) lift(node, nodeToDelete *TreeNode) *TreeNode {
	if node.left != nil {
		node.left = t.lift(node.left, nodeToDelete)
		return node
	} else {
		nodeToDelete.Value = node.Value
		return node.right
	}
}
func (t *TreeNode) Delete(val int, node *TreeNode) *TreeNode {
	// Base case: we've hit bottom of tree. Parent
	// node has no children.
	if node == nil {
		return nil

		// If the value we're deleting is less than or greater
		// than the current node's value, we set the left or right
		// child to be the return value of this recursive call.
	} else if val < node.Value {
		node.left = t.Delete(val, node.left)
		return node
	} else if val > node.Value {
		node.right = t.Delete(val, node.right)
		return node
	} else if val == node.Value {
		// If the current node has no left child, we delete it by returning its
		// right child (and subtree if exists) to be the parent's new subtree.
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			// The current node has two children, we delete the current node
			// by calling the lift function which changes the current node's value to the value
			// of its successor node:
			node.right = t.lift(node.right, node)
			return node
		}
	} else {
		return nil
	}
}

func (t *TreeNode) FindMax(curNode *TreeNode, maxSoFar int) int {
	if curNode == nil {
		return maxSoFar
	}
	if curNode.Value > maxSoFar {
		maxSoFar = curNode.Value
	}
	return t.FindMax(curNode.right, maxSoFar)
}

func main() {
}
