// Double linked list

package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
	curr *Node
}

type Node struct {
	val  string
	next *Node
	prev *Node
}


func (ll *LinkedList) getCurrent() *Node{
	if ll.head != nil {
		return ll.curr  // either nil or a Node ref
	}
	return nil
}

func (ll *LinkedList) getNext() *Node{
	if ll.head == nil {  // empty ll
		return nil
	}
	if ll.curr == nil {  // first time
		ll.curr = ll.head
		return ll.curr
	} else if ll.curr != nil && ll.curr.next != nil { // if there is next one
		ll.curr = ll.curr.next
		return ll.curr
	}
	return nil
}

func (ll *LinkedList) getPrev() *Node{
	if ll.head == nil {
		return nil
	}
	if ll.curr == nil {  // first time
		ll.curr = ll.head
		return ll.curr
	} else if ll.curr != nil && ll.curr.next != nil { // 
		ll.curr = ll.curr.next
		return ll.curr
	}
	return nil
}

func (ll *LinkedList) print(message string) {
	fmt.Println(message, "lengths:", ll.len())
	if ll.head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	curr := ll.head
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next
	}
	fmt.Println()
}

func (ll *LinkedList) addRight(node *Node) {
	curr := (*ll).head
	// iterate until next node is nil
	for curr.next != nil {
		curr = curr.next
	}
	// add prev and next
	curr.next = node
	curr.next.prev = curr
}

func (ll *LinkedList) addLeft(node *Node) {
	curr := (*ll).head
	// reassign the head
	curr.prev = node
	node.next = curr
	ll.head = node
}

func (ll *LinkedList) getByValue(val string) *Node {
	curr := ll.head
	for {
		if curr.val == val {
			return curr
		} else if curr.next == nil {
			return nil
		}
		curr = curr.next
	}
}

func (ll *LinkedList) deleteByValue(val string) {
	curr := ll.head
	for curr != nil {
		if curr.val == val {	
			if curr.prev == nil && curr.next == nil {  // if only
				ll.head = nil
			} else if curr.prev == nil {  // if first
				ll.head = curr.next
				curr.next.prev = nil
			} else if curr.next == nil {  // if last
				curr.prev.next = nil
			} else {  // if in the middle
				curr.prev.next = curr.next
				curr.next.prev = curr.prev.next
			}
		}
		curr = curr.next
	}
}

func (ll *LinkedList) len() int {
	length := 0
	curr := ll.head
	if curr == nil {
		return 0
	}
	for curr != nil {
		length++
		curr = curr.next
	}
	return length
}

func main() {
	var nodes []Node
	nodes = []Node{
		Node{val: "Zero"},
		Node{val: "One"},
	}
	fmt.Printf("%v\n", nodes)

	// Adding nodes to the Linked List
	list := LinkedList{head: &nodes[0]}
	fmt.Printf("Head: %v\n", list.head)
	for i := 0; i < len(nodes); i++ {
		if &nodes[i] == list.head {
			continue
		}
		list.addRight(&nodes[i]) // (*head).Add == head.Add
	}

	rightNode := Node{val: "Two"}
	list.addRight(&rightNode)
	list.print("After right add:")

	newNode := Node{val: "minusOne"}
	list.addLeft(&newNode)
	list.print("After left add:")

	taken := list.getByValue("minusOne")
	fmt.Printf("getByValue: %v\n\n", taken.val)

	list.deleteByValue("minusOne")
	list.print("After deleting first:")
	
	list.deleteByValue("Two")
	list.print("After deleting last:")

	list.addRight(&Node{val: "Two"})
	list.print("Before deleting middle:")
	list.deleteByValue("One")
	list.print("After deleting middle:")

	list.deleteByValue("Zero")
	list.print("Before deleting only (zero is taken out)")
	list.deleteByValue("Two")
	list.print("After deleting only")
	
}
