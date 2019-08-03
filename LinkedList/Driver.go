package main

import "fmt"

func main() {
	ll := &IntSLList{}
	ll.addToHead(8)
	ll.addToHead(7)
	ll.addToHead(7)
	ll.addToHead(7)
	ll.addToTail(20)
	ll.printAll()
	ll.removeDuplicates()
	fmt.Println(ll.isEmpty())
	ll.printAll()
}
