package main

func main() {
}
func sumFormula(max int) int {
	return (max * (max + 1) / 2)
}
func sumLoop(max int) int {
	total := 0
	for val := 0; val < max; val++ {
		total += val
	}
	return total
}
