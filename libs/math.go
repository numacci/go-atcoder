package libs

// Math Utils
func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func nextPermutation(x []int) bool {
	i := len(x) - 2
	if i < 0 {
		return false
	}

	for ; x[i] >= x[i+1]; i-- {
		// completed
		if i == 0 {
			return false
		}
	}

	j := len(x) - 1
	for x[i] >= x[j] {
		j--
	}
	// swap
	x[i], x[j] = x[j], x[i]

	// reverse
	l, r := i+1, len(x)-1
	for l < r {
		x[l], x[r] = x[r], x[l]
		l++
		r--
	}
	return true
}
