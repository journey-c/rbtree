package rbtree

import (
	"fmt"
	"testing"
)

func cmp(a interface{}, b interface{}) int {
	if a.(int) > b.(int) {
		return 1
	} else if a.(int) < b.(int) {
		return -1
	} else {
		return 0
	}
}

var (
	colorRed   = "\x1b[1;31m%v\x1b[0m\n"
	colorBlack = "\x1b[1;34m%v\x1b[0m\n"

	colorCC   = "\033[34m"
	colorLink = "\033[34;1m"
	colorSrc  = "\033[33m"
	colorBin  = "\033[37;1m"
	colorMake = "\033[32;1m"
	colorEnd  = "\033[0m"
)

func printTree(rb *RbTree, node *RbNode, layer int) {
	if node == rb.null {
		return
	}

	printTree(rb, node.right, layer+1)

	for i := 0; i < layer; i++ {
		print("\t")
	}
	if node.color == red {
		fmt.Printf(colorRed, node.K)
	} else {
		fmt.Printf(colorBlack, node.K)
	}

	printTree(rb, node.left, layer+1)
	return
}

func TestRbTree(t *testing.T) {
	tree := NewRbTree(cmp, false)

	keyList := []int{10, 3, 17, 7, 1, 0, 10, 12, 4, 5}

	fmt.Printf(colorMake+"=========== INSERT: %v ==========="+colorEnd+"\n", keyList)
	for _, key := range keyList {
		tree.Insert(key, key*2)
	}
	fmt.Println("level:\t1\t2\t3\t4\t5\t6\t7\t8\t...")
	printTree(tree, tree.root, 1)

	for i := len(keyList) - 1; i >= 0; i-- {
		fmt.Printf(colorMake+"=========== DELETE: %d ==========="+colorEnd+"\n", keyList[i])
		tree.DeleteByKey(keyList[i])
		fmt.Println("level:\t1\t2\t3\t4\t5\t6\t7\t8\t...")
		printTree(tree, tree.root, 1)
	}
}
