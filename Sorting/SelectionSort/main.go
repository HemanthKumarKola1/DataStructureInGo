package main

import "fmt"

func main() {
	input := []int{21, 12, 33, 4, 15, 6}
	fmt.Println(len(input))
	for i := 0; i < len(input); i++ {
		min_idx := i
		for j := i + 1; j < len(input); j++ {
			if input[min_idx] > input[j] {
				min_idx = j
			}
		}
		if min_idx != i {
			y := input[min_idx]
			input[min_idx] = input[i]
			input[i] = y
		}
	}
	fmt.Println(input)
}
