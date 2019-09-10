package main

import "fmt"

func simple(key string) int {
	fmt.Printf("The key is = %s \n", key)
	var len = len(key)
	var sum = 0

	for i := 0; i < len; i++ {
		var decVal = key[i]
		sum = sum + int(decVal)
		fmt.Printf("Character %s = %v \n", string(key[i]), decVal)
	}
	return sum
}

func hashRoutine(keyToSearch string) int {
	var sum = 0
	for _, subSc := range keyToSearch {
		var decVal = keyToSearch[subSc]
		sum += decVal
	}
	return sum
}

func main() {
	var index = simple("1007")
	fmt.Printf("Index number generated from key = %v \n", index)
}
