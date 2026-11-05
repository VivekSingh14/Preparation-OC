package main

import "fmt"

type ListNode struct {
	data     int
	NextNode *ListNode
}

func (ll *ListNode) appendNode(ele int) {
	newNode := &ListNode{data: ele, NextNode: nil}

	temp := ll

	for temp.NextNode != nil {
		temp = temp.NextNode
	}

	temp.NextNode = newNode
}

func (ll *ListNode) displayList() {

	// to ignore 0 otherwise  tempList := ll
	tempList := ll.NextNode
	for tempList != nil {
		fmt.Print(tempList.data, " ")
		tempList = tempList.NextNode
	}
	fmt.Println()
}

func (ll *ListNode) displayMiddleNode() {

	slow := ll.NextNode
	fast := ll.NextNode
	for fast != nil && fast.NextNode != nil {
		slow = slow.NextNode
		fast = fast.NextNode.NextNode
	}
	fmt.Println(slow.data)
}

func main2() {
	ll := ListNode{}

	ll.appendNode(1)
	ll.appendNode(2)
	ll.appendNode(3)
	ll.appendNode(4)
	ll.appendNode(5)
	ll.appendNode(6)
	ll.appendNode(7)
	ll.appendNode(8)
	ll.appendNode(9)

	ll.displayList()

	ll.displayMiddleNode()
}
