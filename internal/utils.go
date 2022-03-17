package internal

import "math/rand"

var MIN = -50
var MAX = 50

func Reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func RandMass(count int) []int {
	var arr []int
	for i := 0; i < count; i++ {
		arr = append(arr, rand.Intn(MAX-MIN)+MIN)
	}
	return arr
}
