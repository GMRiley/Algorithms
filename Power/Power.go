package main

func powerRecurs(base int, exp int) int {
	if exp > 1 {
		return base * powerRecurs(base, exp-1)
	}
	return base
}
func powerLoop(base int, exp int) int {
	var total = base
	for counter := 1; counter < exp; counter++ {
		total *= base
	}
	return total
}
