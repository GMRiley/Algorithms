package main

import (
	"fmt"
)

func main() {
	var array = [8][8]string{
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "Boss", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
	}
	nestedBossFinder(array)
}

//The asymptotic complexity of a nested for loop is O(n^2)
func nestedBossFinder(array [8][8]string) {
	for subX := 0; subX < 8; subX++ {
		for subY := 0; subY < 8; subY++ {
			if array[subX][subY] == "Boss" {
				fmt.Println("Boss found at ", subX, subY)
			}
		}
	}
}
