package main

import (
	"fmt"
)

func merge(arr []int, low int, mid int, high int) []int {

	temp := make([]int, len(arr))
	left := low
	right := mid + 1
	i := 0
	for left <= mid && right <= high {
		if arr[left] < arr[right] {
			temp[i] = arr[left]
			left++
		} else {
			temp[i] = arr[right]
			right++
		}
		i++
	}
	for left <= mid {
		temp[i] = arr[left]
		i++
		left++
	}
	for right <= high {
		temp[i] = arr[right]
		i++
		right++
	}
	i = 0
	for low <= high {
		arr[low] = temp[i]
		i++
		low++
	}
	return arr
}
func mergeSort(arr []int, low int, high int) []int {
	if low >= high {
		return arr
	}
	mid := low + (high-low)/2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	return merge(arr, low, mid, high)
}
func main() {
	arr := []int{323, 2, 25, 325, 20, 34, 5, 2, 55}
	fmt.Println(mergeSort(arr, 0, len(arr)-1))

}
