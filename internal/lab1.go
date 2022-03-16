package internal

import (
	"fmt"
	"math/rand"
	"time"
)

var MIN = -50
var MAX = 50

func BubbleSort(ar []int) {
	var countPerm int
	var countComp int
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			countComp++
			if ar[i] > ar[j] {
				countPerm++
				swap(ar, i, j)
			}
		}
	}
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
}

func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}

func pivot(arr []int, begin, end, countPerm, countComp int) (int, int, int) {
	end--
	temp := arr[begin]
	for begin < end {
		for begin < end {
			countComp++
			if arr[end] < temp {
				arr[begin] = arr[end]
				countPerm++
				begin++
				break
			}
			end--
		}
		for begin < end {
			countComp++
			if arr[begin] > temp {
				arr[end] = arr[begin]
				countPerm++
				end--
				break
			}
			begin++
		}
	}
	arr[begin] = temp
	return begin, countPerm, countComp
}

func RecursiveQuickSort(arr []int, begin, end, countPerm, countComp int) (int, int) {
	if end-begin < 2 {
		return countPerm, countComp
	}
	index, countPerm, countComp := pivot(arr, begin, end, countPerm, countComp)

	countPerm, countComp = RecursiveQuickSort(arr, begin, index, countPerm, countComp)
	countPerm, countComp = RecursiveQuickSort(arr, index+1, end, countPerm, countComp)
	return countPerm, countComp
}

func Lab1() {
	rand.Seed(time.Now().UnixNano())
	ar10 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ar20 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	ar40 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	ranAr10 := randMass(10)
	ranAr20 := randMass(20)
	ranAr40 := randMass(40)

	fmt.Println("------------ BubbleSort ------------ (count: 10)")
	fmt.Println("_____SINGLE_____")
	BubbleSort(ar10)
	Reverse(ar10)
	fmt.Println("_____REVERSE_____")
	BubbleSort(ar10)
	fmt.Println("_____RANDOM_____")

	arr10 := make([]int, 10)
	copy(arr10, ranAr10)
	BubbleSort(arr10)

	fmt.Println("------------ BubbleSort ------------ (count: 20)")
	fmt.Println("_____SINGLE_____")
	BubbleSort(ar20)
	Reverse(ar20)
	fmt.Println("_____REVERSE_____")
	BubbleSort(ar20)
	fmt.Println("_____RANDOM_____")
	arr20 := make([]int, 20)
	copy(arr20, ranAr20)
	BubbleSort(arr20)

	fmt.Println("------------ BubbleSort ------------ (count: 40)")
	fmt.Println("_____SINGLE_____")
	BubbleSort(ar40)
	Reverse(ar40)
	fmt.Println("_____REVERSE_____")
	BubbleSort(ar40)
	fmt.Println("_____RANDOM_____")
	arr40 := make([]int, 40)
	copy(arr40, ranAr40)
	BubbleSort(arr40)

	fmt.Println("------------ QuickSort ------------ (count: 10)")
	fmt.Println("_____SINGLE_____")
	countPerm, countComp := RecursiveQuickSort(ar10, 0, len(ar10), 0, 0)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	Reverse(ar10)
	fmt.Println("_____REVERSE_____")
	countPerm, countComp = RecursiveQuickSort(ar10, 0, len(ar10), 0, 0)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	fmt.Println("_____RANDOM_____")
	copy(arr10, ranAr10)
	countPerm, countComp = RecursiveQuickSort(arr10, 0, len(arr10), 0, 0)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)

	fmt.Println("\n\n------------ QuickSort ------------ (count: 20)")
	fmt.Println("_____SINGLE_____")
	countPerm, countComp = RecursiveQuickSort(ar20, 0, len(ar20), 0, 0)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	Reverse(ar20)
	fmt.Println("_____REVERSE_____")
	countPerm, countComp = RecursiveQuickSort(ar20, 0, len(ar20), 0, 0)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	fmt.Println("_____RANDOM_____")
	copy(arr20, ranAr20)
	countPerm, countComp = RecursiveQuickSort(arr20, 0, len(arr20), 0, 0)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)

	fmt.Println("\n\n------------ QuickSort ------------ (count: 40)")
	fmt.Println("_____SINGLE_____")
	countPerm, countComp = RecursiveQuickSort(ar40, 0, len(ar40), 0, 0)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	Reverse(ar40)
	fmt.Println("_____REVERSE_____")
	countPerm, countComp = RecursiveQuickSort(ar40, 0, len(ar40), 0, 0)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	fmt.Println("_____RANDOM_____")
	copy(arr40, ranAr40)
	countPerm, countComp = RecursiveQuickSort(arr40, 0, len(arr40), 0, 0)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)

}

func randMass(count int) []int {
	var arr []int
	for i := 0; i < count; i++ {
		arr = append(arr, rand.Intn(MAX-MIN)+MIN)
	}
	return arr
}
