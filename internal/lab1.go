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
	fmt.Printf("BubbleSort Permutations: %d\n", countPerm)
	fmt.Printf("BubbleSort Сomparison: %d\n", countComp)
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
