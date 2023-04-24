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

func main() {
	nodeA := Node{data: "a"}
	nodeB := Node{data: "b"}
	nodeC := Node{data: "c"}

	nodeA.nextNode = &nodeB
	nodeB.nextNode = &nodeC
	ll := LinkedList{&nodeA}
	value := ll.Read(1)
	if value != nil {
		fmt.Printf("Value is: %s\n", *value)
	} else {
		fmt.Println("Value is nill")
	}
	value = ll.Read(10)
	if value != nil {
		fmt.Printf("Value is: %s\n", *value)
	} else {
		fmt.Println("Value is nill")
	}
	foundIndex := ll.IndexOf("c")
	fmt.Printf("index of 'c' is: %d\n", *foundIndex)
	foundIndex = ll.IndexOf("d")
	if foundIndex == nil {
		fmt.Println("index of 'd' is nil")
	}

}
