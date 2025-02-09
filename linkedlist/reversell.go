package main

import (
	"fmt"
	"math"
)

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head *node
}

func (l *linkedlist) display() {
	if l.head == nil {
		return
	}

	if l.head.next == nil {
		fmt.Println(l.head.data)
		return
	}

	ptr := l.head

	for ptr != nil {
		fmt.Print(ptr.data, " ")
		ptr = ptr.next
	}
	fmt.Println()
}

func (l *linkedlist) reverse() *node {

	var prev *node
	curr := l.head
	var forward *node

	for curr != nil {
		forward = curr.next
		curr.next = prev
		prev = curr
		curr = forward
	}
	return prev
}

func (l *linkedlist) makecycle() {
	curr := l.head

	var temp *node

	for curr.next != nil {
		if curr.data == 2 {
			temp = curr
		}
		curr = curr.next
	}

	curr.next = temp

}

func (l *linkedlist) hasCycle() bool {

	for l.head.next != nil {
		if l.head.data == math.MaxInt {
			return true
		}
		l.head.data = math.MaxInt
		l.head = l.head.next
	}

	return false
}

func (l *linkedlist) hasCycleOtherMethod() bool {

	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast.data == slow.data {
			return true
		}
	}

	return false
}
func main() {

	ll := linkedlist{}
	ll.insert(1)
	ll.insert(2)
	ll.insert(3)
	ll.insert(4)
	ll.insert(5)
	//ll.insert(6)

	//iterate through the linkedlist
	ll.display()

	//reverse the linkedlist
	ll.head = ll.reverse()
	ll.display()

	ll1 := linkedlist{}
	ll1.insert(1)
	ll1.insert(2)
	ll1.insert(3)
	ll1.insert(4)
	ll1.insert(5)

	ll1.makecycle()

	// check if linked list has loop
	fmt.Println(ll1.hasCycle())

	fmt.Println(ll.hasCycleOtherMethod())

	//merge 2 linkedlists
	ll2 := linkedlist{}

	ll2.insert(1)
	ll2.insert(3)
	ll2.insert(5)
	ll2.insert(7)

	ll2.display()

	ll3 := linkedlist{}

	ll3.insert(2)
	ll3.insert(4)
	ll3.insert(6)
	ll3.insert(8)

	ll3.display()

	//ll4 := linkedlist{}

	//ll4.mergell(ll2, ll3)

	test := mergell(ll2, ll3)

	//ll4.display()

	test.display()

	fmt.Println("-------------------------")

	//merge k linkedlists

	ll4 := linkedlist{}

	ll4.insert(1)
	ll4.insert(4)
	ll4.insert(7)
	ll4.insert(10)

	ll4.display()

	fmt.Println("-------------------------")

	ll5 := linkedlist{}

	ll5.insert(2)
	ll5.insert(5)
	ll5.insert(8)
	ll5.insert(11)

	ll5.display()

	fmt.Println("-------------------------")

	ll6 := linkedlist{}

	ll6.insert(3)
	ll6.insert(6)
	ll6.insert(9)
	ll6.insert(12)

	ll6.display()

	fmt.Println("-------------------------")

	var hr []linkedlist
	hr = append(hr, ll4)
	hr = append(hr, ll5)
	hr = append(hr, ll6)

	tet := mergeKLists(hr)

	tet.display()

	//removed the nth node from end
	removed := removeNthFromEnd(ll6, 2)

	removed.display()

	fmt.Println("-------------------------")

	//reorder list
	ll7 := linkedlist{}
	ll7.insert(1)
	ll7.insert(2)
	ll7.insert(3)
	ll7.insert(4)
	ll7.insert(5)
	//ll.insert(6)

	//iterate through the linkedlist
	ll7.display()

	reor := reorderList(ll7)
	reor.display()

}

// func (ll4 *linkedlist) mergell(l linkedlist, l1 linkedlist) {
func mergell(l linkedlist, l1 linkedlist) *linkedlist {

	temp := l.head
	temp2 := l1.head

	ll4 := linkedlist{}

	if temp.data < temp2.data {
		ll4.insert(temp.data)
		temp = temp.next
	} else if temp.data > temp2.data {
		ll4.insert(temp2.data)
		temp2 = temp2.next
	}

	for temp != nil || temp2 != nil {

		if temp != nil && temp2 != nil {
			if temp.data < temp2.data {
				ll4.insert(temp.data)
				temp = temp.next
			} else {
				ll4.insert(temp2.data)
				temp2 = temp2.next
			}
		} else if temp != nil {
			ll4.insert(temp.data)
			break
		} else {
			ll4.insert(temp2.data)
			break
		}
	}
	return &ll4
}

// merge k linked lists
func mergeKLists(lists []linkedlist) *linkedlist {

	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		l1 := lists[0]
		l2 := lists[1]
		lists = lists[2:]

		merged := mergell(l1, l2)
		lists = append(lists, *merged)
	}
	return &lists[0]
}

func (l *linkedlist) insert(data int) {

	n := node{}
	n.data = data
	if l.head == nil {
		l.head = &n
		return
	}

	ptr := l.head

	for ptr != nil {
		if ptr.next == nil {
			ptr.next = &n
			return
		}
		ptr = ptr.next
	}

}

func removeNthFromEnd(l linkedlist, n int) *linkedlist {
	fast := l.head
	slow := l.head

	for i := 0; i < n; i++ {
		fast = fast.next
	}

	if fast == nil {
		return &l
	}

	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	slow.next = slow.next.next

	return &l
}

func reorderList(l linkedlist) *linkedlist {

	if l.head == nil {
		return &l
	}

	var slice []*node

	temp := l.head

	for temp != nil {
		slice = append(slice, temp)
		temp = temp.next
	}

	low, r := 0, len(slice)-1

	for low < r {
		slice[r].next = nil
		slice[low].next = slice[r]
		low++
		slice[r].next = slice[low]
		slice[low].next = nil
		r--

	}

	l.head = slice[0]

	return &l

}
