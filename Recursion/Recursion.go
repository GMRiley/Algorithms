package main

func reverse(phrase string) string {
	if len(phrase) <= 1 {
		return phrase
	}
	return phrase[len(phrase)-1:] + reverse(phrase[:len(phrase)-1])
}
