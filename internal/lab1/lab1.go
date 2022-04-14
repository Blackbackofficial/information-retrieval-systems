package lab1

import (
	"fmt"
	"math/rand"
	"retrieval-systems/internal"
	"time"
)

// BubbleSort
func bubbleSort(ar []int) {
	var countPerm int
	var countComp int
	for i := 0; i < len(ar)-1; i++ {
		for j := 0; j < len(ar)-i-1; j++ {
			countComp++
			if ar[j] > ar[j+1] {
				countPerm++
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
}

// InsertionSort
func insertionSort(ar []int) {
	var countPerm int
	var countComp int
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		countComp++
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
			countPerm++
		}
		ar[j] = x
	}
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	fmt.Println(ar)
}

// QuickSort
func quickSort(arr []int, begin, end, countPerm, countComp int) (int, int) {
	if end-begin < 2 {
		return countPerm, countComp
	}
	index, countPerm, countComp := pivot(arr, begin, end, countPerm, countComp)

	countPerm, countComp = quickSort(arr, begin, index, countPerm, countComp)
	countPerm, countComp = quickSort(arr, index+1, end, countPerm, countComp)
	//fmt.Println(arr)
	return countPerm, countComp
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

// Lab1 Print output
func Lab1() {
	//arr := rand.Perm(10000)
	//fmt.Println(arr)
	rand.Seed(time.Now().UnixNano())
	ar10 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ar20 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	ar40 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	ranAr10 := internal.RandMass(10)
	ranAr20 := internal.RandMass(20)
	ranAr40 := internal.RandMass(40)

	fmt.Println("------------ BubbleSort ------------ (count: 10)")
	fmt.Println("_____SINGLE_____")
	fmt.Println(ar10)
	bubbleSort(ar10)
	fmt.Println(ar10)
	internal.Reverse(ar10)
	fmt.Println("_____REVERSE_____")
	fmt.Println(ar10)
	bubbleSort(ar10)
	fmt.Println(ar10)
	fmt.Println("_____RANDOM_____")
	arr10 := make([]int, 10)
	copy(arr10, ranAr10)
	fmt.Println(arr10)
	bubbleSort(arr10)
	fmt.Println(arr10)

	fmt.Println("------------ BubbleSort ------------ (count: 20)")
	fmt.Println("_____SINGLE_____")
	fmt.Println(ar20)
	bubbleSort(ar20)
	fmt.Println(ar20)
	internal.Reverse(ar20)
	fmt.Println("_____REVERSE_____")
	fmt.Println(ar20)
	bubbleSort(ar20)
	fmt.Println(ar20)
	fmt.Println("_____RANDOM_____")
	arr20 := make([]int, 20)
	copy(arr20, ranAr20)
	fmt.Println(arr20)
	bubbleSort(arr20)
	fmt.Println(arr20)

	fmt.Println("------------ BubbleSort ------------ (count: 40)")
	fmt.Println("_____SINGLE_____")
	fmt.Println(ar40)
	bubbleSort(ar40)
	fmt.Println(ar40)
	internal.Reverse(ar40)
	fmt.Println("_____REVERSE_____")
	fmt.Println(ar40)
	bubbleSort(ar40)
	fmt.Println(ar40)
	fmt.Println("_____RANDOM_____")
	arr40 := make([]int, 40)
	copy(arr40, ranAr40)
	fmt.Println(arr40)
	bubbleSort(arr40)
	fmt.Println(arr40)

	fmt.Println("\n\n------------ QuickSort ------------ (count: 10)")
	fmt.Println("_____SINGLE_____")
	fmt.Println(ar10)
	countPerm, countComp := quickSort(ar10, 0, len(ar10), 0, 0)
	fmt.Println(ar10)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	internal.Reverse(ar10)
	fmt.Println("_____REVERSE_____")
	fmt.Println(ar10)
	countPerm, countComp = quickSort(ar10, 0, len(ar10), 0, 0)
	fmt.Println(arr10)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	fmt.Println("_____RANDOM_____")
	copy(arr10, ranAr10)
	fmt.Println(arr10)
	countPerm, countComp = quickSort(arr10, 0, len(arr10), 0, 0)
	fmt.Println(arr10)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)

	fmt.Println("------------ QuickSort ------------ (count: 20)")
	fmt.Println("_____SINGLE_____")
	fmt.Println(ar20)
	countPerm, countComp = quickSort(ar20, 0, len(ar20), 0, 0)
	fmt.Println(ar20)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	internal.Reverse(ar20)
	fmt.Println("_____REVERSE_____")
	fmt.Println(ar20)
	countPerm, countComp = quickSort(ar20, 0, len(ar20), 0, 0)
	fmt.Println(ar20)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	fmt.Println("_____RANDOM_____")
	copy(arr20, ranAr20)
	fmt.Println(arr20)
	countPerm, countComp = quickSort(arr20, 0, len(arr20), 0, 0)
	fmt.Println(arr20)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)

	fmt.Println("------------ QuickSort ------------ (count: 40)")
	fmt.Println("_____SINGLE_____")
	fmt.Println(ar40)
	countPerm, countComp = quickSort(ar40, 0, len(ar40), 0, 0)
	fmt.Println(ar40)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	internal.Reverse(ar40)
	fmt.Println("_____REVERSE_____")
	fmt.Println(ar40)
	countPerm, countComp = quickSort(ar40, 0, len(ar40), 0, 0)
	fmt.Println(ar40)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)
	fmt.Println("_____RANDOM_____")
	copy(arr40, ranAr40)
	fmt.Println(arr40)
	countPerm, countComp = quickSort(arr40, 0, len(arr40), 0, 0)
	fmt.Println(arr40)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)

	fmt.Println("\n\n------------ InsertionSort ------------ (count: 10)")
	fmt.Println("_____SINGLE_____")
	fmt.Println(ar10)
	insertionSort(ar10)
	internal.Reverse(ar10)
	fmt.Println("_____REVERSE_____")
	fmt.Println(ar10)
	insertionSort(ar10)
	fmt.Println(ar10)
	fmt.Println("_____RANDOM_____")
	copy(arr10, ranAr10)
	fmt.Println(arr10)
	insertionSort(arr10)
	fmt.Println(arr10)

	fmt.Println("------------ InsertionSort ------------ (count: 20)")
	fmt.Println("_____SINGLE_____")
	fmt.Println(ar20)
	insertionSort(ar20)
	fmt.Println(ar20)
	internal.Reverse(ar20)
	fmt.Println("_____REVERSE_____")
	fmt.Println(ar20)
	insertionSort(ar20)
	fmt.Println(ar20)
	fmt.Println("_____RANDOM_____")
	copy(arr20, ranAr20)
	fmt.Println(arr20)
	insertionSort(arr20)
	fmt.Println(arr20)

	fmt.Println("------------ InsertionSort ------------ (count: 40)")
	fmt.Println("_____SINGLE_____")
	fmt.Println(ar40)
	insertionSort(ar40)
	internal.Reverse(ar40)
	fmt.Println("_____REVERSE_____")
	fmt.Println(ar40)
	insertionSort(ar40)
	fmt.Println("_____RANDOM_____")
	copy(arr40, ranAr40)
	fmt.Println(arr40)
	insertionSort(arr40)
}
