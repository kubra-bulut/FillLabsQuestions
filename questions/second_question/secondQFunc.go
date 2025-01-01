package second_question

import "fmt"

// GenerateOutput recursively generates the sequence.
func GenerateOutput(n int) {
	if n < 2 {
		return
	}
	next := n / 2
	GenerateOutput(next)
	fmt.Println(n)
}
