package main

import (
	"fmt"
	"math"
	"time"
)

// Algorithm from https://stackoverflow.com/a/21854246

// Only primes less than or equal to N will be generated
func allPrimes(N int) []int {

	var x, y, n int
	nsqrt := math.Sqrt(float64(N))

	is_prime := make([]bool, N)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= N && (n%12 == 1 || n%12 == 5) {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) + y*y
			if n <= N && n%12 == 7 {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= N && n%12 == 11 {
				is_prime[n] = !is_prime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if is_prime[n] {
			for y = n * n; y < N; y += n * n {
				is_prime[y] = false
			}
		}
	}

	is_prime[2] = true
	is_prime[3] = true

	primes := make([]int, 0, 1270606)
	for x = 0; x < len(is_prime)-1; x++ {
		if is_prime[x] {
			primes = append(primes, x)
		}
	}

	// primes is now a slice that contains all primes numbers up to N
	return primes
}

func bloat(mb int) string {
	b := make([]byte, mb*1024*1024)
	b[0] = 1
	b[len(b)-1] = 1
	return fmt.Sprintf("Allocated %v Mb of memory.", mb)
}

func prime(max int) string {
	p := allPrimes(max)
	if len(p) > 0 {
		return fmt.Sprintf("The largest prime less than %v is %v.", max, p[len(p)-1])
	} else {
		return fmt.Sprintf("There are no primes smaller than %v.", max)
	}
}

func sleep(ms int) string {
	start := time.Now().UnixNano()
	time.Sleep(time.Duration(ms) * time.Millisecond)
	end := time.Now().UnixNano()
	return fmt.Sprintf("Slept for %.2f milliseconds.", float64(end-start)/1000000)
}
