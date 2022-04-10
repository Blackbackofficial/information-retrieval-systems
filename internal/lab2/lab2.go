package lab2

import "fmt"

// MergeSort
func mergeSort(items []int, countPerm, countComp int) ([]int, int, int) {
	if len(items) < 2 {
		return items, countPerm, countComp
	}
	first, countPerm, countComp := mergeSort(items[:len(items)/2], countPerm, countComp)
	second, countPerm, countComp := mergeSort(items[len(items)/2:], countPerm, countComp)
	return merge(first, second, countPerm, countComp)
}

func merge(a []int, b []int, countPerm, countComp int) ([]int, int, int) {
	var final []int
	var i, j int
	for i < len(a) && j < len(b) {
		countComp++
		if a[i] < b[j] {
			final = append(final, a[i])
			countPerm++
			i++
		} else {
			final = append(final, b[j])
			countPerm++
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
		countPerm++
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
		countPerm++
	}
	return final, countPerm, countComp
}

func Lab2() {
	arr := []int{54, 26, 93, 17, 77, 31, 44, 55, 20, 40, 20, 55, 1, 7, 39, 11, 25, 19, 34, 10}
	fmt.Println("------------ Balanced MergeSort ------------ (count: 20)")
	arr, countPerm, countComp := mergeSort(arr, 0, 0)
	fmt.Println(arr)
	fmt.Printf("Permutations: %d\n", countPerm)
	fmt.Printf("Сomparison: %d\n", countComp)

	fmt.Println("------------ Oscillated Sorting (MergeSort) ------------ (count: 20)")
}
