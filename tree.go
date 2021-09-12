package main

import (
	"fmt"
)

type Tree struct {
	Data      int
	LeftNode  *Tree
	RightNode *Tree
}

func createTree() *Tree {
	var root *Tree = &Tree{4, nil, nil}
	root.LeftNode = &Tree{2, nil, nil}
	root.LeftNode.LeftNode = &Tree{1, nil, nil}
	root.LeftNode.RightNode = &Tree{3, nil, nil}
	root.RightNode = &Tree{6, nil, nil}
	root.RightNode.LeftNode = &Tree{5, nil, nil}
	root.RightNode.RightNode = &Tree{7, nil, nil}
	return root
}
func preOrderTraverseTree(root *Tree) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	preOrderTraverseTree(root.LeftNode)
	preOrderTraverseTree(root.RightNode)
}
func inOrderTraverseTree(root *Tree) {
	if root == nil {
		return
	}

	inOrderTraverseTree(root.LeftNode)
	fmt.Println(root.Data)
	inOrderTraverseTree(root.RightNode)
}
func postOrderTraverseTree(root *Tree) {
	if root == nil {
		return
	}
	postOrderTraverseTree(root.LeftNode)
	postOrderTraverseTree(root.RightNode)
	fmt.Println(root.Data)
}
func bfs(root *Tree) {
	if root == nil {
		return
	}
	// for root 需要藉助隊列
	var queue []*Tree
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		fmt.Println(node.Data)
		if node.LeftNode != nil {
			queue = append(queue, node.LeftNode)
		}
		if node.RightNode != nil {
			queue = append(queue, node.RightNode)
		}
		queue = queue[1:] // 通過這樣的方式達到出隊列
	}
}
func dfsInverseTree(root *Tree) {
	if root == nil {
		return
	}
	// 交換左右子樹
	tempNode := root.LeftNode
	root.LeftNode = root.RightNode
	root.RightNode = tempNode
	dfsInverseTree(root.LeftNode)
	// fmt.Println(root.Data)
	dfsInverseTree(root.RightNode)
}

func bfsInverseTree(root *Tree) {
	if root == nil {
		return
	}
	// for root 需要藉助隊列
	var queue []*Tree
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		// fmt.Println(node.Data)
		// 交換左右子樹
		tempNode := node.LeftNode
		node.LeftNode = node.RightNode
		node.RightNode = tempNode
		if node.LeftNode != nil {
			queue = append(queue, node.LeftNode)
		}
		if node.RightNode != nil {
			queue = append(queue, node.RightNode)
		}
		queue = queue[1:] // 通過這樣的方式達到出隊列
	}
}

func levelTraverseTree(root *Tree) [][]*Tree {
	if root == nil {
		return [][]*Tree{}
	}
	var arr [][]*Tree
	var oneArr []*Tree
	var tempArr []*Tree
	oneArr = append(oneArr, root)
	arr = append(arr, oneArr)
	for i := 0; i < len(oneArr); {
		node := oneArr[i]
		if node.LeftNode != nil {
			tempArr = append(tempArr, node.LeftNode)
		}
		if node.RightNode != nil {
			tempArr = append(tempArr, node.RightNode)
		}
		if i == len(oneArr)-1 && len(tempArr) > 0 {
			oneArr = tempArr
			arr = append(arr, tempArr)
			tempArr = []*Tree{}
			i = 0
			continue
		}
		i++
	}
	return arr
}

func main() {
	root := createTree()
	fmt.Println("pre-------")
	preOrderTraverseTree(root)
	fmt.Println("in-------")
	inOrderTraverseTree(root)
	fmt.Println("post------")
	postOrderTraverseTree(root)
	fmt.Println("bfs------")
	bfs(root)
	fmt.Println("defInverseTree-----")
	dfsInverseTree(root)
	fmt.Println("inorder-----")
	inOrderTraverseTree(root)
	fmt.Println("bfsInverseTree------")
	bfsInverseTree(root)
	fmt.Println("inorder---------")
	inOrderTraverseTree(root)
}
