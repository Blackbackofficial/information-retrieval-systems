package internal

import "fmt"

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
