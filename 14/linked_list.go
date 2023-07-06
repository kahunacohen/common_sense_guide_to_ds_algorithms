package main

import "fmt"

// A linked list implementation.
type Node struct {
	data     string
	nextNode *Node
}
type LinkedList struct {
	FirstNode *Node
}

func (l LinkedList) Read(index int) *string {
	currentNode := l.FirstNode
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

func (l LinkedList) IndexOf(value string) *int {
	var curIndex int
	curNode := l.FirstNode
	for curNode.data != value {
		curIndex += 1
		curNode = curNode.nextNode
		if curNode == nil {
			return nil
		}
	}
	return &curIndex
}

func (l LinkedList) InsertAtIndex(indx int, value string) {
	newNode := Node{data: value, nextNode: nil}
	if indx == 0 {
		newNode.nextNode = l.FirstNode
	} else {
		curIndx := 0
		curNode := l.FirstNode
		for curIndx < indx-1 {
			curNode = curNode.nextNode
			curIndx += 1
		}
		newNode.nextNode = curNode.nextNode
		curNode.nextNode = &newNode
	}
}

func (l *LinkedList) DeleteAtIndex(indx int) {
	if indx == 0 {
		l.FirstNode = l.FirstNode.nextNode
	}
}

func main() {
	fmt.Println("Hello world")
}
