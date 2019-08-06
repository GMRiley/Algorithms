package main

import "fmt"

type IntSLList struct {
	head *IntSLLNode
	tail *IntSLLNode
}

func (p *IntSLList) isEmpty() bool {
	return p.head == nil
}
func (p *IntSLList) addToHead(el int) {
	p.head = &IntSLLNode{el, p.head}
	if p.tail == nil {
		p.tail = p.head
	}
}
func (p *IntSLList) addToTail(el int) {
	if !p.isEmpty() {
		p.tail.next = &IntSLLNode{el, nil}
		p.tail = p.tail.next
	} else {
		p.tail = &IntSLLNode{el, nil}
		p.head = p.tail
	}
}
func (p *IntSLList) deleteFromHead() int {
	el := p.head.info
	if p.head == p.tail {
		p.tail = nil
		p.head = p.tail
	} else {
		p.head = p.head.next
	}
	return el
}
func (p *IntSLList) deleteFromTail() int {
	el := p.tail.info
	if p.head == p.tail {
		p.tail = nil
		p.head = p.tail
	} else {
		tmp := &IntSLLNode{el, nil}
		for tmp = p.head; tmp.next != p.tail; tmp = tmp.next {
		}
		p.tail = tmp
		p.tail.next = nil
	}
	return el
}
func (p *IntSLList) printAll() {
	for tmp := p.head; tmp != nil; tmp = tmp.next {
		fmt.Println(tmp.info, " ")
	}
}
func (p *IntSLList) isInList(el int) bool {
	tmp := p.head
	for tmp != nil && tmp.info != el {
		tmp = tmp.next
	}
	return tmp != nil
}
func (p *IntSLList) delete(el int) {
	if !p.isEmpty() {
		if p.head == p.tail && el == p.head.info {
			p.tail = nil
			p.head = p.tail
		} else if el == p.head.info {
			p.head = p.head.next
		} else {
			pred, tmp := p.head, p.head.next
			for tmp != nil && tmp.info != el {
				pred, tmp = pred.next, tmp.next
			}
			if tmp != nil {
				pred.next = tmp.next
				if tmp == p.tail {
					p.tail = pred
				}
			}
		}
	}
}

func (p *IntSLList) sum() int {
	sum := 0
	for tmp := p.head; tmp != nil; tmp = tmp.next {
		sum += tmp.info
	}
	return sum
}
func (p *IntSLList) max() int {
	max := p.head.info
	for tmp := p.head; tmp != nil; tmp = tmp.next {
		if tmp.info > max {
			max = tmp.info
		}
	}
	return max
}
func (p *IntSLList) min() int {
	min := p.head.info
	for tmp := p.head; tmp != nil; tmp = tmp.next {
		if tmp.info < min {
			min = tmp.info
		}
	}
	return min
}
func (p *IntSLList) removeDuplicates() {
	pred, curr := p.head, p.head
	integers := make(map[int]bool)

	for curr != nil {
		var curVal = curr.info
		if integers[curVal] {
			pred.next = curr.next
		} else {
			integers[curr.info] = true
			pred = curr
		}
		curr = curr.next
	}
}
func (p *IntSLList) removeMedian() {
	count := 0
	for tmp := p.head; tmp != nil; tmp = tmp.next {
		count++
	}
	count = count / 2
	tmpTwo := p.head
	for counter := 0; counter < count; counter++ {
		tmpTwo = tmpTwo.next
	}
	val := tmpTwo.info
	p.delete(val)
}
