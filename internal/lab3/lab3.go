package lab3

// Index returns the first index substr found in the s.
func Index(s string, substr string) int {
	d := CalculateSlideTable(substr)
	return IndexWithTable(&d, s, substr)
}

// IndexWithTable returns the first index substr found in the s.
func IndexWithTable(d *[256]int, s string, substr string) int {
	lSub := len(substr)
	ls := len(s)
	// fmt.Println(ls, lSub)
	switch {
	case lSub == 0:
		return 0
	case lSub > ls:
		return -1
	case lSub == ls:
		if s != substr {
			return -1
		}
		return 0
	}

	i := 0
	for i+lSub-1 < ls {
		j := lSub - 1
		if j >= 0 && s[i+j] == substr[j] {
			j--
		}
		if j < 0 {
			return i
		}

		slid := j - d[s[i+j]]
		if slid < 1 {
			slid = 1
		}
		i += slid
	}
	return -1
}

// CalculateSlideTable builds sliding amount per each unique byte in the substring
func CalculateSlideTable(substr string) [256]int {
	var d [256]int
	for i := 0; i < 256; i++ {
		d[i] = -1
	}
	for i := 0; i < len(substr); i++ {
		d[substr[i]] = i
	}
	return d
}

func Lab3() {

}