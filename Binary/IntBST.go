package main

import (
	"fmt"
	"io"
)

type IntBST struct {
	root *IntBSTNode
	prev *IntBSTNode
}

func (p *IntBST) visit() {
	fmt.Println(p.root.key)
}
func (p *IntBST) insert(el int) {
	for p.root != nil {
		p.prev = p.root
		if p.root.key < el {
			p.root = p.root.right
		} else {
			p.root = p.root.left
		}
	}

	if p.root == nil {
		p.root.key = el
	} else if p.prev.key < el {
		p.prev.right.key = el
	} else {
		p.prev.left.key = el
	}
}

func (p *IntBST) recursInsert(el int) *IntBST {
	if p.root == nil {
		p.root = &IntBSTNode{key: el, left: nil, right: nil}
	} else {
		p.root.recursInsert(el)
	}
	return p
}
func (n *IntBSTNode) recursInsert(el int) {
	if n == nil {
		return
	} else if el <= n.key {
		if n.left == nil {
			n.left = &IntBSTNode{key: el, left: nil, right: nil}
		} else {
			n.left.recursInsert(el)
		}
	} else {
		if n.right == nil {
			n.right = &IntBSTNode{key: el, left: nil, right: nil}
		} else {
			n.right.recursInsert(el)
		}
	}
}
func print(w io.Writer, node *IntBSTNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.key)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}
func (p *IntBST) search(el int) *IntBSTNode {
	for p.root != nil {
		if el == p.root.key {
			return p.root
		} else if el < p.root.key {
			p.root = p.root.left
		} else {
			p.root = p.root.right
		}
	}
	return nil
}
func (p *IntBST) DFS(el int) {
	fmt.Print("Depth First Traversal: ")
	var traversal = func(n *IntBSTNode) {
		if n == nil {
			return
		}

	}
}
