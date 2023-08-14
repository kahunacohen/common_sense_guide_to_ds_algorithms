package main

import "testing"

func TestSearch(t *testing.T) {
	left := TreeNode{Value: 0}
	right := TreeNode{Value: 2}
	root := TreeNode{Value: 1, left: &left, right: &right}

	// Search for a value that doens't exist.
	res := root.Search(3)
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
