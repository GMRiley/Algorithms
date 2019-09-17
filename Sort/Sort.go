package main

type employee struct {
	age     int
	name    string
	address string
}

func BubbleSort(employees []employee) []employee {
	for i := len(employees); i > 0; i-- {
		for j := 1; j < i; j++ {
			if employees[j-1].name > employees[j].name {
				intermediate := employees[j]
				employees[j] = employees[j-1]
				employees[j-1] = intermediate
			}
		}
	}
	return employees
}
