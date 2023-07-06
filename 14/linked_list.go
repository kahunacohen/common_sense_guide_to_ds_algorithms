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

func main() {
	fmt.Println("Hello world")
}

// func main() {
// 	nodeA := Node{data: "a"}
// 	nodeB := Node{data: "b"}
// 	nodeC := Node{data: "c"}

// 	nodeA.nextNode = &nodeB
// 	nodeB.nextNode = &nodeC
// 	ll := LinkedList{&nodeA}
// 	value := ll.Read(1)
// 	if value != nil {
// 		fmt.Printf("Value is: %s\n", *value)
// 	} else {
// 		fmt.Println("Value is nill")
// 	}
// 	value = ll.Read(10)
// 	if value != nil {
// 		fmt.Printf("Value is: %s\n", *value)
// 	} else {
// 		fmt.Println("Value is nill")
// 	}
// 	foundIndex := ll.IndexOf("c")
// 	fmt.Printf("index of 'c' is: %d\n", *foundIndex)
// 	foundIndex = ll.IndexOf("d")
// 	if foundIndex == nil {
// 		fmt.Println("index of 'd' is nil")
// 	}
// 	ll.InsertAtIndex(1, "new node")
// 	fmt.Println(*(ll.Read(0)))
// 	fmt.Println(*(ll.Read(1)))
// 	fmt.Println(*(ll.Read(2)))
// 	fmt.Println(*(ll.Read(3)))
// }
