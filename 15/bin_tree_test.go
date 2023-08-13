package main

import "testing"

func TestSearch(t *testing.T) {
	left := TreeNode{value: 0}
	right := TreeNode{value: 2}
	root := TreeNode{value: 1, left: &left, right: &right}
	//res := root.Search(1, &root)
	// if res.value != 1 {
	// 	t.Fatalf("wanted 1, got %d", res.value)
	// }
	// res = root.Search(2, &root)
	// if res.value != 2 {
	// 	t.Fatalf("wanted 2, got %d", res.value)
	// }
	res := root.Search(3, &root)
	if res.value != 3 {
		t.Fatalf("wanted 3, got %d", res.value)
	}

}
