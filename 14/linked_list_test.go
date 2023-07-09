package main

import (
	"testing"
)

func getLinkedList() LinkedList {
	n1 := Node{data: "a"}
	n2 := Node{data: "b"}
	n3 := Node{data: "c"}
	n1.nextNode = &n2
	n2.nextNode = &n3
	n3.nextNode = nil
	linkedList := LinkedList{&n1}
	return linkedList
}

func TestRead(t *testing.T) {
	linkedList := getLinkedList()
	val := linkedList.Read(1)
	if *val != "b" {
		t.Fatalf("wanted 'b', got %s", *val)
	}
}

func TestIndexOf(t *testing.T) {
	linkedList := getLinkedList()
	indx := linkedList.IndexOf("c")
	if *indx != 2 {
		t.Fatalf("wanted 2, got %d", *indx)
	}
}

func TestInsertAtIndex(t *testing.T) {
	linkedList := getLinkedList()
	linkedList.InsertAtIndex(3, "d")
	val := linkedList.Read(3)
	if *val != "d" {
		t.Fatalf("wanted 'd', got %s", *val)
	}
}

func TestDeleteAtIndex0(t *testing.T) {
	linkedList := getLinkedList()
	linkedList.DeleteAtIndex(0)
	val := linkedList.Read(0)
	if *val != "b" {
		t.Fatalf("wanted 'b', got '%s'", *val)
	}
}

func TestDeleteAtLastIndex(t *testing.T) {
	linkedList := getLinkedList()
	linkedList.DeleteAtIndex(2)
	val := linkedList.Read(0)
	if *val != "a" {
		t.Fatalf("wanted 'a', got '%s'", *val)
	}
	val = linkedList.Read(1)
	if *val != "b" {
		t.Fatalf("wanted 'b', got '%s'", *val)
	}
	val = linkedList.Read(2)
	if val != nil {
		t.Fatalf("wanted nil, didn't get it")
	}
}

func TestDeleteAtIndex(t *testing.T) {
	linkedList := getLinkedList()
	linkedList.DeleteAtIndex(1)
	val := linkedList.Read(0)
	if *val != "a" {
		t.Fatalf("wanted 'a', got '%s'", *val)
	}
	val = linkedList.Read(1)
	if *val != "c" {
		t.Fatalf("wanted 'c', got '%s'", *val)
	}
}
