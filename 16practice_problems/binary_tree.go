package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

/*
         1
       /   \
      2     3
    /   \ /   \
   4    5 6    7
  /
 9

*/

func Contructor(value int) *Node {
	return &Node{value, nil, nil}
}

func preorder(root *Node, arr *[]int) {
	if root == nil {
		return
	}

	*arr = append(*arr, root.value)

	preorder(root.left, arr)
	preorder(root.right, arr)
}

func inorder(root *Node, arr *[]int) {
	if root == nil {
		return
	}
	inorder(root.left, arr)
	*arr = append(*arr, root.value)
	inorder(root.right, arr)
}

func postorder(root *Node, arr *[]int) {
	if root == nil {
		return
	}
	postorder(root.left, arr)
	postorder(root.right, arr)
	*arr = append(*arr, root.value)
}

func main() {
	root := Contructor(1)

	root.left = Contructor(2)
	root.right = Contructor(3)

	root.left.left = Contructor(4)
	root.left.right = Contructor(5)
	root.right.left = Contructor(6)
	root.right.right = Contructor(7)

	root.left.left.left = Contructor(9)

	arr := []int{}

	preorder(root, &arr)
	fmt.Println(arr)

	arr = []int{}

	inorder(root, &arr)
	fmt.Println(arr)
}
