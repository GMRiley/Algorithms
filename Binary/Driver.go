package main

import (
	"fmt"
	"os"
)

func main() {
	myTree := IntBST{}
	myTree.recursInsert(13)
	myTree.recursInsert(10)
	myTree.recursInsert(2)
	myTree.recursInsert(12)
	myTree.recursInsert(25)
	myTree.recursInsert(20)
	myTree.recursInsert(31)
	myTree.recursInsert(29)
	fmt.Println(myTree.search(13))
	print(os.Stdout, myTree.root, 0, 'M')
	fmt.Println(myTree.DFS(13, myTree.root))
}
