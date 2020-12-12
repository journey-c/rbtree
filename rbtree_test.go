package rbtree

import (
	"fmt"
	"math/rand"
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
	tree := NewRbTree(cmp, true)

	var keyList []int
	for i := 0; i < 1000000; i++ {
		key := rand.Int() % 10
		// fmt.Printf(colorMake+"=========== INSERT: %v ==========="+colorEnd+"\n", key)
		keyList = append(keyList, key)
		tree.Insert(key, key)
	}
	// fmt.Println("level:\t1\t2\t3\t4\t5\t6\t7\t8\t...")
	// printTree(tree, tree.root, 1)

	for i := len(keyList) - 1; i >= 0; i-- {
		// fmt.Printf(colorMake+"=========== DELETE: %d ==========="+colorEnd+"\n", keyList[i])
		tree.DeleteByKey(keyList[i])
		// fmt.Println("level:\t1\t2\t3\t4\t5\t6\t7\t8\t...")
		// fmt.Println("Size:", tree.Size())
		// printTree(tree, tree.root, 1)
	}
	printTree(tree, tree.root, 1)
	tree.DeleteByKey(11)
}
