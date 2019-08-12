package main

type data struct {
	array []int
}

func (d *data) clear() {
	d.array = nil
}

func (d *data) push(element int) {
	d.array = append(d.array, element)
}

func (d *data) pop() {
	length := len(d.array)
	d.array = d.array[:length-1]
}
