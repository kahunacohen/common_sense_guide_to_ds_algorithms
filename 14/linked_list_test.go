package main

import (
	"testing"
)

func getLinkedList() LinkedList {
	n1 := Node{Data: "a"}
	n2 := Node{Data: "b"}
	n3 := Node{Data: "c"}
	n1.NextNode = &n2
	n2.NextNode = &n3
	n3.NextNode = nil
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

func TestInsertAtEnd(t *testing.T) {
	dl := DoublyLinkedList{}
	dl.InsertAtEnd("end")
	if dl.FirstNode.data != "end" && dl.LastNode.data == "end" {
		t.Fatalf("wanted 'end', got first: %s, last: %s", dl.FirstNode.data, dl.LastNode.data)
	}
	dl.InsertAtEnd("new end")
	if dl.LastNode.data != "new end" {
		t.Fatalf("wanted 'new end', got %s", dl.LastNode.data)
	}
	if dl.LastNode.prevNode.data != "end" {
		t.Fatalf("wanted 'end', got %s", dl.LastNode.prevNode.data)
	}
}

func TestGetValues(t *testing.T) {
	ll := getLinkedList()
	res := ll.GetValues()
	if len(res) != 3 {
		t.Fatalf("wanted 3, got %d", len(res))
	}
	if res[0] != "a" && res[1] != "b" && res[2] != "c" {
		t.Fatalf("didn't get a slice with the right string values. Got: %v", res)
	}
}
func TestGetAllValuesFromEnd(t *testing.T) {
	dl := DoublyLinkedList{}
	dl.InsertAtEnd("a")
	dl.InsertAtEnd("b")
	dl.InsertAtEnd("c")
	res := dl.GetAllValuesFromEnd()
	if res[0] != "c" {
		t.Fatalf("wanted 'c', got %s", res[0])
	}
	if res[1] != "b" {
		t.Fatalf("wanted 'b', got %s", res[0])
	}
	if res[2] != "a" {
		t.Fatalf("wanted 'a', got %s", res[0])
	}
}

func TestGetLastNode(t *testing.T) {
	ll := getLinkedList()
	lastNode := ll.GetLastNode()
	if lastNode == nil {
		t.Fatalf("last node shouldn't be nil")
	}
	if lastNode.Data != "c" {
		t.Fatalf("wanted 'c', got %s", lastNode.Data)
	}
}

func TestReverseListIterative(t *testing.T) {
	ll := getLinkedList()
	ll.ReverseIterative()
	vals := ll.GetValues()
	c := vals[0]
	b := vals[1]
	a := vals[2]
	if c != "c" {
		t.Fatalf("wamted 'c', got %s", c)
	}
	if b != "b" {
		t.Fatalf("wamted 'b', got %s", b)
	}
	if a != "a" {
		t.Fatalf("wamted 'a', got %s", a)
	}
}
