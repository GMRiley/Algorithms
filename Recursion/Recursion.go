package main

import (
	"strconv"
)

func reverse(phrase string) string {
	if len(phrase) <= 1 {
		return phrase
	}
	return phrase[len(phrase)-1:] + reverse(phrase[:len(phrase)-1])
}

func say(phrase string, iter int) string {
	if iter == 0 {
		return phrase
	}
	counter := 0
	holder := ""
	holderVal := rune(phrase[0])
	for _, nums := range phrase {
		if nums == holderVal {
			counter++
		} else {
			holder = holder + strconv.Itoa(counter) + string(holderVal)
			holderVal = nums
			counter = 1
		}
	}
	holder = holder + strconv.Itoa(counter) + string(holderVal)
	return say(holder, iter-1)
}
