//This is binary search tree also

package main

import (
	"fmt"
	"math"
)

type node struct {
	left  *node
	data  int
	right *node
}

type bst struct {
	root *node
}

func (t *bst) InsertIntoTree(data int) {
	if t.root == nil {
		t.root = &node{data: data}
	} else {
		t.root.InsertIntoTree(data)
	}
}

func (nod *node) InsertIntoTree(data int) {
	if data <= nod.data {
		if nod.left == nil {
			nod.left = &node{data: data}
		} else {
			nod.left.InsertIntoTree(data)
		}
	} else {
		if nod.right == nil {
			nod.right = &node{data: data}
		} else {
			nod.right.InsertIntoTree(data)
		}
	}
}
func printInOrder(n *node) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%d \t", n.data)
		printInOrder(n.right)
	}
}

func main1() {

	//insert into bst or bt
	var t bst
	t.InsertIntoTree(4)
	t.InsertIntoTree(2)
	t.InsertIntoTree(5)
	t.InsertIntoTree(1)
	t.InsertIntoTree(3)
	t.InsertIntoTree(10)
	//t.InsertIntoTree(8)
	//t.InsertIntoTree(6)
	//t.InsertIntoTree(9)
	//t.InsertIntoTree(7)

	//print in order using recursion
	printInOrder(t.root)

	fmt.Println()

	//max depth or height of bt/bst
	fmt.Println(maxDepth(t))

	var t1 bst
	t1.InsertIntoTree(4)
	t1.InsertIntoTree(2)
	t1.InsertIntoTree(5)
	t1.InsertIntoTree(1)
	t1.InsertIntoTree(3)
	//t.InsertIntoTree(10)
	//t.InsertIntoTree(8)
	//t.InsertIntoTree(6)
	//t.InsertIntoTree(9)
	//t.InsertIntoTree(7)

	printInOrder(t1.root)

	fmt.Println()

	//to check if 2 bt/bsts are same
	fmt.Println(sametrees(t.root, t1.root))

	//Invert bt/bst
	var t2 bst
	t2.InsertIntoTree(4)
	t2.InsertIntoTree(2)
	t2.InsertIntoTree(5)
	t2.InsertIntoTree(1)
	t2.InsertIntoTree(3)

	printInOrder(t2.root)

	fmt.Println()

	// t2 = invertBT(t2)

	// printInOrder(t2.root)

	//level order traversal
	levelOrderTraversal(t2)

	fmt.Println()

	var t3 bst

	t3.InsertIntoTree(2)
	t3.InsertIntoTree(1)
	t3.InsertIntoTree(3)

	printInOrder(t3.root)

	fmt.Println()

	// check if given t3 is subtree of t2
	fmt.Println(isSubTree(t2, t3))

	fmt.Println()

	//to valiate if given tree is bst, we are iterating it inorder traverse so that we get the data in increasing order.
	fmt.Println(validateBST(t2))

	fmt.Println()

	//Same concept as above for finding kth smallest element
	//to valiate if given tree is bst, we are iterating it inorder traverse so that we get the data in increasing order.
	fmt.Println(kthSmallestElement(t2, 2))

}

func maxDepth(n bst) int {

	if n.root == nil {
		return 0
	}
	var queue []*node
	temp := n.root
	queue = append(queue, temp)
	queue = append(queue, nil)
	depth := 0

	for len(queue) != 0 {
		temp = queue[0]
		queue = queue[1:]
		if temp == nil {
			depth++
		}
		if temp != nil {
			if temp.left != nil {
				queue = append(queue, temp.left)
			}

			if temp.right != nil {
				queue = append(queue, temp.right)
			}
		} else if len(queue) != 0 {
			queue = append(queue, nil)
		}
	}

	return depth
}

func sametrees(t *node, t1 *node) bool {
	if t == nil && t1 == nil {
		return true
	}

	if t == nil || t1 == nil {
		return false
	}

	return sametrees(t.left, t1.left) && sametrees(t.right, t1.right)

}

func invertBT(t bst) bst {

	if t.root == nil {
		fmt.Println("No node in the tree")
		return t
	}
	var queue []*node
	temp := t.root

	queue = append(queue, temp)

	for len(queue) != 0 {
		node1 := queue[0]
		test := node1.left
		node1.left = node1.right
		node1.right = test
		queue = queue[1:]

		if node1.left != nil {
			queue = append(queue, node1.left)
		}
		if node1.right != nil {
			queue = append(queue, node1.right)
		}
	}
	return t
}

func levelOrderTraversal(t bst) {
	if t.root == nil {
		fmt.Println("tree without root node")
		return
	}

	var queue []*node
	temp := t.root
	queue = append(queue, temp)
	//queue = append(queue, nil)

	for len(queue) != 0 {
		test := queue[0]

		fmt.Print(test.data, " ")

		if test.left != nil {
			queue = append(queue, test.left)
		}

		if test.right != nil {
			queue = append(queue, test.right)

		}
		queue = queue[1:]
	}

}

func isSubTree(t1 bst, t2 bst) bool {

	var queue []*node
	queue = append(queue, t1.root)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.data == t2.root.data && sametrees(curr, t2.root) {
			return true
		}

		if curr.left != nil {
			queue = append(queue, curr.left)
		}

		if curr.right != nil {
			queue = append(queue, curr.right)

		}

	}
	return false

}

func validateBST(t bst) bool {
	if t.root == nil {
		return true
	}

	temp := t.root

	var stack []*node

	prev := math.MinInt

	for len(stack) != 0 || temp != nil {
		if temp != nil {
			stack = append(stack, temp)
			temp = temp.left
		} else {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop.data <= prev {
				return false
			}
			prev = pop.data
			temp = pop.right
		}
	}
	return true
}

func kthSmallestElement(t bst, k int) int {
	if t.root == nil {
		return -1
	}

	temp := t.root
	var stack []*node
	stack = append(stack, temp)
	temp = temp.left
	count := 1

	for len(stack) != 0 || temp != nil {
		if temp != nil {
			stack = append(stack, temp)
			temp = temp.left
		} else {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if count == k {
				return pop.data
			}
			temp = pop.right
			count = count + 1
		}
	}
	return 0
}
