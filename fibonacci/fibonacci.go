package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	next := 0
	prev := 0
	cur := 0
	return func() int {
		prev = cur
		cur = next
		if cur == 0 {
			next = 1
		} else {
			next = cur + prev
		}
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
