package main

import "fmt"

func main() {
	input := []int{21, 12, 33, 4, 15, 6}
	fmt.Println(len(input))
	for i := len(input) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if input[j] > input[j+1] {
				y := input[j]
				input[j] = input[j+1]
				input[j+1] = y
			}
		}
	}
	fmt.Println(input)
}
