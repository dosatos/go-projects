// Double linked list

package main

import (
	"fmt"
)

type Node struct {
	val  string
	next *Node
	prev *Node
}

func (curr *Node) Add(node *Node) {
	// iterates until the end of the LL
	// to add new Node
	for {
		if curr.next == nil {
			curr.next, (*node).prev = node, curr // node.prev == (*node).prev
			break
		}
		curr = curr.next
	}
}

func main() {
	var nodes []Node
	nodes = []Node{
		Node{val: "Zero"},
		Node{val: "One"},
		Node{val: "Two"},
	}
	fmt.Printf("%v\n", nodes)

	// Adding nodes to the Linked List
	head := &nodes[0]
	for i := 0; i < len(nodes); i++ {
		if &nodes[i] == head {
			continue
		}
		(*head).Add(&nodes[i]) // (*head).Add == head.Add
	}
	fmt.Printf("%v\n", nodes)
}
