package lab3

import "fmt"

// NaiveString naive string search algorithm
func naiveString(src, subStr string) (out []int) {
	n, m := len(src), len(subStr)

	for i := 0; i <= n-m; i++ {
		if subStr == src[i:i+m] {
			out = append(out, i)
		}
	}

	return out
}

// BoyerMoore Boyer-Moore string search algorithm
func boyerMoore(src, subStr string) (out []int) {
	lenT, lenP := len(src), len(subStr)
	s := 0

	letters := uniqueLetters(src)

	occ, bPos, shift := make(map[uint8]int, len(letters)), make([]int, lenP+1), make([]int, lenP+1)

	for _, letter := range letters {
		occ[letter] = -1
	}

	l := len(subStr) // init occurrence
	for j := 0; j < l; j++ {
		occ[subStr[j]] = j
	}
	preprocessStrongSuffix(bPos, shift, subStr)

	j := bPos[0] // preprocess border
	for i := 0; i <= l; i++ {
		if shift[i] == 0 {
			shift[i] = j
		}
		if i == j {
			j = bPos[j]
		}
	}

	for s <= lenT-lenP {
		j = lenP - 1
		for ; j >= 0 && subStr[j] == src[s+j]; j-- {
		}

		if j < 0 {
			out = append(out, s)
			s += shift[0]
		} else {
			s += max(shift[j+1], j-occ[src[s+j]])
		}
	}
	return out
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func uniqueLetters(s string) []uint8 {
	tmp := make([]uint8, 0, len(s)/2)
	for j := range s {
		flag := false
		for i := 0; i < len(tmp) && !flag; i++ {
			flag = tmp[i] == s[j]
		}
		if !flag {
			tmp = append(tmp, s[j])
		}
	}
	return tmp
}

func preprocessStrongSuffix(bPos, shift []int, p string) {
	m := len(p)
	i, j := m, m+1
	bPos[i] = j

	for i > 0 {
		for j <= m && p[i-1] != p[j-1] {
			if shift[j] == 0 {
				shift[j] = j - i
			}

			j = bPos[j]
		}
		i--
		j--
		bPos[i] = j
	}
}

// KnuthMorrisPrat Knuth-Morris-Pratt string search algorithm
func knuthMorrisPrat(src, subStr string) (out []int) {
	n, m := len(src), len(subStr)
	pi := computePrefixFunction(subStr)
	q := -1

	for i := 0; i < n; i++ {
		for q > -1 && subStr[q+1] != src[i] {
			q = pi[q]
		}

		if subStr[q+1] == src[i] {
			q++
		}

		if q == m-1 {
			out = append(out, i-m+1)
			q = pi[q]
		}
	}
	return out
}

func computePrefixFunction(subStr string) []int {
	m := len(subStr)
	pi := make([]int, m)
	k := -1
	pi[0] = -1
	for q := 1; q < m; q++ {
		for k > -1 && subStr[k] != subStr[q] {
			k = pi[k]
		}
		if subStr[k+1] == subStr[q] {
			k++
		}
		pi[q] = k
	}
	return pi
}

func Lab3() {
	src := `A cryptographic approach to secure information implies its transformation which enables it to be read
only by the owner of the secret key. The reliability of a cryptographic method of securing data depends`
	subStr := `approach to secure`
	fmt.Printf("Naive string search: %s, %s, %v\n", src, subStr, naiveString(src, subStr))
	fmt.Printf("Boyer-Moore string search: %s, %s, %v\n", src, subStr, boyerMoore(src, subStr))
	fmt.Printf("KnuthMorrisPrat string search: %s, %s, %v\n", src, subStr, knuthMorrisPrat(src, subStr))
}
