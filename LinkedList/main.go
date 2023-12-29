package main

import "fmt"

type Node struct {
	data     int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(data int) {
	newHead := Node{
		data: data,
	}
	if linkedList.headNode != nil {
		newHead.nextNode = linkedList.headNode
	}
	linkedList.headNode = &newHead
}

func (linkedList *LinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.data)
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return nil
}

func (linkedList *LinkedList) AddToEnd(data int) {
	node := &Node{data: data}
	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(4)
	linkedList.AddToHead(3)
	linkedList.IterateList()
	fmt.Println(linkedList.LastNode().data)
}
