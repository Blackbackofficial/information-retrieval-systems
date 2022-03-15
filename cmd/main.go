package main

import (
	"fmt"
	"retrieval-systems/internal"
)

func main() {
	ar10 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ar20 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	ar40 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	internal.BubbleSort(ar10)
	internal.BubbleSort(ar20)
	internal.BubbleSort(ar40)
	internal.BubbleSort(ar10)
	internal.Reverse(ar10)
	internal.BubbleSort(ar10)

	fmt.Println(ar10)

}
