package lab2

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
