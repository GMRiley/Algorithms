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

func (d *data) pop() int {
	length := len(d.array) - 1
	holder := d.array[length]
	d.array = d.array[:length]
	return holder
}
func (d *data) size() int {
	return len(d.array)
}
func (d *data) find(num int) bool {
	for _, v := range d.array {
		if v == num {
			return true
		}
	}
	return false
}
func (d *data) indexOf(element int) int {
	for i, v := range d.array {
		if v == element {
			return i
		}
	}
	return -1
}
