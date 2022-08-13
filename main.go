package main

import (
	"fmt"

	"gitlab.com/wshaman/hw-concurrency/lib/batch"
)

func main() {
	b1 := batch.GetBatch(10, 1)
	b2 := batch.GetBatch(10, 1)
	fmt.Println(b1, "\n", b2)
	fmt.Println(equal(b1, b2))
}

// equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []batch.User) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
