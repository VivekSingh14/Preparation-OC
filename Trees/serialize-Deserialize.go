package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec struct for serialization/deserialization
type Codec struct {
}

// serialize converts a binary tree into a string
func (tmp *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil,"
	}
	return strconv.Itoa(root.Val) + "," + tmp.serialize(root.Left) + tmp.serialize(root.Right)
}

// deserializeList is a standalone function that deserializes a list of strings into a binary tree
func deserializeList(strs *[]string) *TreeNode {
	if len(*strs) == 0 || (*strs)[0] == "nil" {
		*strs = (*strs)[1:]
		return nil
	}
	num, _ := strconv.Atoi((*strs)[0])
	*strs = (*strs)[1:]
	root := &TreeNode{Val: num}
	root.Left = deserializeList(strs)
	root.Right = deserializeList(strs)
	return root
}

// deserialize converts a string representation into a binary tree
func (tmp *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	strs := strings.Split(data, ",")
	return deserializeList(&strs)
}

// printTree prints the tree in a pre-order traversal for verification
func printTree(node *TreeNode) {
	if node == nil {
		fmt.Print("nil ")
		return
	}
	fmt.Printf("%d ", node.Val)
	printTree(node.Left)
	printTree(node.Right)
}

func main() {
	codec := &Codec{}
	// Example input: represents a binary tree [1,2,3,nil,nil,4,5]
	input := "1,2,nil,nil,3,4,nil,nil,5,nil,nil"
	root := codec.deserialize(input)
	fmt.Println("Deserialized tree (pre-order traversal):")
	printTree(root)
	fmt.Println()
	// Serialize the tree back to a string
	serialized := codec.serialize(root)
	fmt.Println("Serialized tree:", serialized)
}
