package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func insertAtTail(head *Node, value int) {
	LNode := Node{value, nil}

	if head == nil {
		head = &LNode
		return
	}

	var temp *Node = head

	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = &LNode
}

func display(head *Node) {

	var temp *Node = head

	for temp != nil {
		fmt.Printf("%v ", temp.Data)

		temp = temp.Next
	}
	fmt.Println()

}

func search_number(head *Node, key int) bool {

	var temp *Node = head

	for temp != nil {
		if temp.Data == key {
			return true
		}
		temp = temp.Next
	}
	return false
}

func reverseList(head *Node) *Node {

	if head == nil {
		return head
	}

	var prev *Node = nil
	var cur *Node = head
	var temp *Node

	for cur != nil {

		temp = cur.Next

		cur.Next = prev

		prev = cur

		cur = temp
	}

	return prev
}

func main() {
	head := Node{} // 0 is inserted here

	insertAtTail(&head, 1)
	insertAtTail(&head, 2)
	insertAtTail(&head, 3)
	insertAtTail(&head, 4)

	display(&head)

	key := 5

	if search_number(&head, key) {
		fmt.Println("Given key is found")
	} else {
		fmt.Println("Given key is not found")
	}

	reverseList(&head)
	display(&head)
}
