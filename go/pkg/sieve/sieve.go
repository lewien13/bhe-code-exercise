package sieve

import (
	"math"
)

type Sieve interface {
	NthPrime(n int64) int64
}

func NewSieve() Sieve {
	s := new(sieve)
	s.primes = []int64{2, 3}
	return s
}

type sieve struct {
	primes []int64
}

// NthPrime returns the nth prime in the zero indexed series all prime numbers
func (s sieve) NthPrime(n int64) int64 {
	//fails for negative values of n, but requirements
	// don't allow for error returns

	//find the prime if we have already calculated it
	lastIndex := len(s.primes) - 1
	if lastIndex >= int(n) {
		return s.primes[n]
	}

	//overestimate the upper bound for nth prime
	m := math.Log(float64(n+1)) * float64(n+1) * 2
	maximum := int64(math.Max(m, 3))

	//create a bool array to use as the sieve
	possibilities := make([]bool, maximum)

	//mark composites of already calculated primes
	for i := 1; i <= lastIndex; i++ {
		p := s.primes[i]
		for comp := p * p; comp <= int64(maximum)-1; comp = comp + p + p {
			possibilities[comp] = true
		}
	}

	//calculate new primes
	p := s.primes[lastIndex]
	for p = p + 2; p <= int64(math.Ceil(math.Sqrt(float64(maximum)))); p = p + 2 {
		//if p is already marked as composite continue
		if possibilities[p] {
			continue
		}

		//append to primes slice
		s.primes = append(s.primes, p)

		//mark odd multiples of p as composite starting with p^2
		for comp := p * p; comp < maximum; comp = comp + p + p {
			possibilities[comp] = true
		}
	}

	//add any non-composite number between sqrt(maximum) and maximum to the prime slice
	for ; p < maximum; p = p + 2 {
		if !possibilities[p] {
			s.primes = append(s.primes, p)
		}
	}

	return s.primes[n]
}
