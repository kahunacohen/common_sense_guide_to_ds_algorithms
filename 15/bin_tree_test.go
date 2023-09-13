package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	left := TreeNode{Value: 0}
	right := TreeNode{Value: 2}
	root := TreeNode{Value: 1, left: &left, right: &right}

	// Search for a value that doens't exist.
	res := root.Search(3)
	fmt.Println(res)
	if res != nil {
		t.Fatalf("wanted nil, got %d", res.Value)
	}

	res = root.Search(1)
	if res.Value != 1 {
		t.Fatalf("wanted 1, got %d", res.Value)
	}
	res = root.Search(2)
	if res.Value != 2 {
		t.Fatalf("wanted 2, got %d", res.Value)
	}
}

func TestInsert(t *testing.T) {
	left := TreeNode{Value: 0}
	right := TreeNode{Value: 2}
	root := TreeNode{Value: 1, left: &left, right: &right}
	root.Insert(4)
	fmt.Printf("Root is: %d\n", root.Value)
	fmt.Printf("left of root is: %d\n", root.left.Value)
	fmt.Printf("right of root is: %d\n", root.right.Value)
	fmt.Printf("right of right is: %d\n", right.right.Value)
	if right.right.Value != 4 {
		t.Fatalf("wanted 4, got: %d", right.right.Value)
	}
	root.Insert(3)
	fmt.Printf("\nRoot is: %d\n", root.Value)
	fmt.Printf("left of root is: %d\n", root.left.Value)
	fmt.Printf("right of root is: %d\n", root.right.Value)
	fmt.Printf("right of right is: %d\n", right.right.Value)
	fmt.Printf("left of right of right is: %d\n", right.right.left.Value)
	if right.right.left.Value != 3 {
		t.Fatalf("wanted 3, got: %d", right.right.Value)
	}
	fmt.Println(root.Search(3))
}
func getTree() TreeNode {
	left := TreeNode{Value: 0}
	right := TreeNode{Value: 2}
	root := TreeNode{Value: 1, left: &left, right: &right}
	return root
}
func TestDeleteLeftLeaf(t *testing.T) {
	root := getTree()
	root.Delete(0, &root)
	if root.Value != 1 {
		t.Fatalf("wanted 1, got %d", root.Value)
	}
	if root.left != nil {
		t.Fatalf("wanted nil, got: %v", root.left)
	}
	if root.right.Value != 2 {
		t.Fatalf("wanted 2, got %d", root.right.Value)
	}
}
func TestDeleteRightLeaf(t *testing.T) {
	root := getTree()
	root.Delete(2, &root)
	if root.Value != 1 {
		t.Fatalf("wanted 1, got %d", root.Value)
	}
	if root.left.Value != 0 {
		t.Fatalf("wanted 0, got: %v", root.left)
	}
	if root.right != nil {
		t.Fatalf("wanted nil, got %d", root.right.Value)
	}
}
func TestFindMax(t *testing.T) {
	zero := TreeNode{Value: 0}
	one := TreeNode{Value: 1}
	root := TreeNode{Value: 2}
	three := TreeNode{Value: 3}
	four := TreeNode{Value: 4}
	five := TreeNode{Value: 5}

	root.left = &zero
	root.right = &four
	zero.left = &one
	four.left = &three
	four.right = &five

	max := root.FindMax()
	if max != 5 {
		t.Fatalf("wanted 3, got %d", max)
	}
}
