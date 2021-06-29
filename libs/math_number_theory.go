package libs

import (
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for p := 2; p*p <= n; p++ {
		if n%p == 0 {
			return false
		}
	}
	return true
}

func primeFactorisation(n int) [][2]int {
	ns := int(math.Sqrt(float64(n)))
	res := make([][2]int, 0, ns)

	for p := 2; p*p <= n; p++ {
		if n%p != 0 {
			continue
		}
		exp := 0
		for n%p == 0 {
			exp++
			n /= p
		}
		res = append(res, [2]int{p, exp})
	}

	if n != 1 {
		res = append(res, [2]int{n, 1})
	}
	return res
}

// eratosthenes return the prime numbers below n
func eratosthenes(n int) []int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	res := make([]int, 0, n)
	for i := 2; i < n; i++ {
		if isPrime[i] {
			res = append(res, i)
			// Remove multiples of i
			for j := i * 2; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return res
}
