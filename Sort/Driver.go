package main

import "fmt"

func main() {
	employees := []employee{
		employee{age: 25, name: "Alice", address: "Somewhere Ave"},
		employee{age: 32, name: "Denise", address: "Rainbow Ave"},
		employee{age: 22, name: "Ben", address: "Over Ave"},
		employee{age: 30, name: "Chad", address: "The Ave"},
	}

	employees = BubbleSort(employees)
	fmt.Println(employees)
}
