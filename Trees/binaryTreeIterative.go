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
	t.InsertIntoTree1(6)
	t.InsertIntoTree1(1)
	t.InsertIntoTree1(3)
	t.InsertIntoTree1(5)
	t.InsertIntoTree1(7)

	fmt.Println(preOrder(t.root))

	fmt.Println(inorder(t.root))

}

func preOrder(root *node1) []int {

	var stack []*node1
	var res []int

	curr := root

	stack = append(stack, curr)

	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, temp.data)

		if temp.right != nil {
			stack = append(stack, temp.right)
		}

		if temp.left != nil {
			stack = append(stack, temp.left)
		}
	}
	return res
}

func inorder(root *node1) []int {
	var stack []*node1
	var res []int

	curr := root

	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.data)
		curr = curr.right
	}
	return res

}
