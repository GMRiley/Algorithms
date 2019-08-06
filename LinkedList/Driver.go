package main

import "fmt"

func main() {
	ll := &IntSLList{}
	ll.addToHead(8)
	ll.addToHead(8)
	ll.addToHead(8)
	ll.addToHead(8)
	ll.addToHead(7)
	ll.addToHead(7)
	ll.addToHead(7)
	ll.addToTail(20)
	fmt.Println("Printing current values")
	ll.printAll()
	ll.removeDuplicates()
	fmt.Println("printing values after duplicate removal")
	ll.printAll()
	ll.removeMedian()
	fmt.Println("printing after median removal")
	ll.printAll()
}
