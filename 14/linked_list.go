package main

import "fmt"

// A linked list implementation.
type Node struct {
	Data     string
	NextNode *Node
}
type LinkedList struct {
	FirstNode *Node
}

func (l LinkedList) Read(index int) *string {
	currentNode := l.FirstNode
	currentIndex := 0
	// Keep looking for the next node until index is the same
	for currentIndex < index {
		currentNode = currentNode.NextNode
		currentIndex += 1
		if currentNode == nil {
			return nil
		}

	}
	return &currentNode.Data
}

func (l LinkedList) GetValues() []string {
	var res []string
	currentNode := l.FirstNode
	for currentNode != nil {
		res = append(res, currentNode.Data)
		currentNode = currentNode.NextNode
	}
	return res
}

func (l LinkedList) IndexOf(value string) *int {
	var curIndex int
	curNode := l.FirstNode
	for curNode.Data != value {
		curIndex += 1
		curNode = curNode.NextNode
		if curNode == nil {
			return nil
		}
	}
	return &curIndex
}

func (l LinkedList) InsertAtIndex(indx int, value string) {
	newNode := Node{Data: value, NextNode: nil}
	if indx == 0 {
		newNode.NextNode = l.FirstNode
	} else {
		curIndx := 0
		curNode := l.FirstNode
		for curIndx < indx-1 {
			curNode = curNode.NextNode
			curIndx += 1
		}
		newNode.NextNode = curNode.NextNode
		curNode.NextNode = &newNode
	}
}

func (l *LinkedList) DeleteAtIndex(indx int) {
	if indx == 0 {
		l.FirstNode = l.FirstNode.NextNode
	}
	curNode := l.FirstNode
	curIndx := 0
	// Find the node preceeding the one we want to delete.
	for curIndx < indx-1 {
		curNode = curNode.NextNode
		curIndx++
	}
	fmt.Println(curIndx)
	fmt.Println(curNode.Data)
	curNode.NextNode = curNode.NextNode.NextNode
}
func (l *LinkedList) GetLastNode() *Node {
	if l.FirstNode == nil {
		return nil
	}
	curNode := l.FirstNode
	for {
		if curNode.NextNode == nil {
			return curNode
		}
		curNode = curNode.NextNode
	}
}
func (l *LinkedList) ReverseIterative() {
	var prev, cur, next *Node
	prev = nil
	cur = l.FirstNode
	for cur != nil {
		next = cur.NextNode
		cur.NextNode = prev
		prev = cur
		cur = next
	}
	l.FirstNode = prev
}

// This is an intellectual excersise that is not given in the
// problem set. In real life it doesn't make sense to write this
// recursively as it could lead to stack overflow with linked lists
// with lots of elements. Even if it doesn't it uses O(n) auxilary
// memory space. This is because each function call uses up a stack frame.
// In a language with tail optimization such as Haskell, OCaml, Scala etc.
// this would make more sense.
func (l *LinkedList) ReverseRecursive() {
	var prev *Node

	// Call a private method to start things off. Otherwise
	// it's awkward to pass initial parameters to
	// the public method.
	l.FirstNode = l.swapNodes(l.FirstNode, prev)
}

func (l *LinkedList) swapNodes(cur, prev *Node) *Node {
	if cur == nil {
		return prev
	}
	next := cur.NextNode
	cur.NextNode = prev
	prev = cur
	cur = next
	return l.swapNodes(cur, prev)
}

// Given a node in a linked list, with no access to the linked
// list itself, delete the node
func DeleteTheNode(node *Node) {
	// first copy the data and pointer of this node.
	node.Data = node.NextNode.Data
	node.NextNode = node.NextNode.NextNode
	fmt.Println("foo")
}

type DoublyLinkedNode struct {
	data     string
	nextNode *DoublyLinkedNode
	prevNode *DoublyLinkedNode
}

type DoublyLinkedList struct {
	FirstNode *DoublyLinkedNode
	LastNode  *DoublyLinkedNode
}

func (dl DoublyLinkedList) Read(index int) *string {
	currentNode := dl.FirstNode
	currentIndex := 0
	// Keep looking for the next node until index is the same
	for currentIndex < index {
		currentNode = currentNode.nextNode
		currentIndex += 1
		if currentNode == nil {
			return nil
		}

	}
	return &currentNode.data
}
func (dl *DoublyLinkedList) InsertAtEnd(value string) {
	// No nodes yet.
	node := &DoublyLinkedNode{data: value, nextNode: nil, prevNode: nil}
	if dl.FirstNode == nil {
		dl.FirstNode = node
		dl.LastNode = node
	} else {
		node.prevNode = dl.LastNode
		dl.LastNode.nextNode = node
		dl.LastNode = node
	}
}
func (dl *DoublyLinkedList) GetAllValuesFromEnd() []string {
	if dl.FirstNode.nextNode == nil {
		return []string{dl.FirstNode.data}
	}
	var ret []string
	currentNode := dl.LastNode
	for currentNode != nil {
		ret = append(ret, currentNode.data)
		currentNode = currentNode.prevNode
	}
	return ret
}

func main() {
	fmt.Println("Hello world")
}
