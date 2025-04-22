package main

import "fmt"

type node1 struct {
	left  *node1
	data  int
	right *node1
}

type bst1 struct {
	root *node1
}

func (t *bst1) InsertIntoTree1(data int) {
	if t.root == nil {
		t.root = &node1{data: data}
	} else {
		t.root.InsertIntoTree1(data)
	}
}

func (nod *node1) InsertIntoTree1(data int) {
	if data <= nod.data {
		if nod.left == nil {
			nod.left = &node1{data: data}
		} else {
			nod.left.InsertIntoTree1(data)
		}
	} else {
		if nod.right == nil {
			nod.right = &node1{data: data}
		} else {
			nod.right.InsertIntoTree1(data)
		}
	}
}

func main() {

	var t bst1
	t.InsertIntoTree1(4)
	t.InsertIntoTree1(2)
	t.InsertIntoTree1(5)
	t.InsertIntoTree1(1)
	t.InsertIntoTree1(3)
	t.InsertIntoTree1(10)

	fmt.Println(preOrder(t.root))

}

func preOrder(root *node1) []int {

	var stack []*node1
	var res []int

	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			res = append(res, curr.data)
			stack = append(stack, curr.left)
			curr = curr.left
		} else {
			stack = stack[:len(stack)-1]
			curr = stack[len(stack)-1]
			curr = curr.right

		}
	}
	return res
}
