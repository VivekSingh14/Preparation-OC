package main

import "fmt"

type ListNode1 struct {
	data     int
	NextNode *ListNode1
}

func (ll *ListNode1) appendNode(ele int) {
	newNode := &ListNode1{data: ele, NextNode: nil}

	temp := ll

	for temp.NextNode != nil {
		temp = temp.NextNode
	}

	temp.NextNode = newNode
}

func (ll *ListNode1) displayList() {

	// to ignore 0 otherwise  tempList := ll
	tempList := ll.NextNode
	for tempList != nil {
		fmt.Print(tempList.data, " ")
		tempList = tempList.NextNode
	}
	fmt.Println()
}

func (ll *ListNode1) removeDuplicateNode() {

	slow := ll.NextNode
	fast := ll.NextNode.NextNode
	for fast != nil {
		if fast.data == slow.data {
			fast = fast.NextNode
		} else {
			slow.NextNode = fast
			slow = slow.NextNode
			fast = fast.NextNode
		}

	}
	slow.NextNode = nil
}

func main4() {
	ll := ListNode1{}

	ll.appendNode(1)
	ll.appendNode(1)
	ll.appendNode(2)
	ll.appendNode(3)
	ll.appendNode(3)

	ll.displayList()

	ll.removeDuplicateNode()

	ll.displayList()
}
