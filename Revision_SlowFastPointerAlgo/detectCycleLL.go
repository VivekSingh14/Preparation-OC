package main

import "fmt"

type ListNode2 struct {
	data     int
	NextNode *ListNode2
}

func (ll *ListNode2) appendNode(ele int) {
	newNode := &ListNode2{data: ele, NextNode: nil}

	temp := ll

	for temp.NextNode != nil {
		temp = temp.NextNode
	}

	temp.NextNode = newNode
}

func (ll *ListNode2) makingCycle() {
	temp := ll.NextNode

	temp1 := ll

	for temp1.NextNode != nil {
		temp1 = temp1.NextNode
	}
	temp1.NextNode = temp
}

func (ll *ListNode2) displayList() {

	// to ignore 0 otherwise  tempList := ll
	tempList := ll.NextNode
	for tempList != nil {
		fmt.Print(tempList.data, " ")
		tempList = tempList.NextNode
	}
	fmt.Println()
}

func (ll *ListNode2) checkCycleInLL() bool {

	slow := ll.NextNode
	fast := ll.NextNode
	for fast != nil && fast.NextNode != nil {
		slow = slow.NextNode
		fast = fast.NextNode.NextNode
		if slow == fast {
			return true
		}

	}
	return false
}

func main() {
	ll := ListNode2{}

	ll.appendNode(1)
	ll.appendNode(2)
	ll.appendNode(3)
	ll.appendNode(4)
	ll.appendNode(5)
	ll.appendNode(6)

	ll.displayList()

	//intentionally to test the logic
	//ll.makingCycle()

	fmt.Println(ll.checkCycleInLL())

}
