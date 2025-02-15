package sieve

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthPrime(t *testing.T) {
	sieve := NewSieve()

	assert.Equal(t, int64(2), sieve.NthPrime(0))
	assert.Equal(t, int64(3), sieve.NthPrime(1))
	assert.Equal(t, int64(5), sieve.NthPrime(2))
	assert.Equal(t, int64(7), sieve.NthPrime(3))
	assert.Equal(t, int64(23), sieve.NthPrime(8))
	assert.Equal(t, int64(71), sieve.NthPrime(19))
	assert.Equal(t, int64(541), sieve.NthPrime(99))
	assert.Equal(t, int64(3581), sieve.NthPrime(500))
	assert.Equal(t, int64(7793), sieve.NthPrime(986))
	assert.Equal(t, int64(17393), sieve.NthPrime(2000))
	assert.Equal(t, int64(1299721), sieve.NthPrime(100000))
	assert.Equal(t, int64(15485867), sieve.NthPrime(1000000))
	assert.Equal(t, int64(179424691), sieve.NthPrime(10000000))
	assert.Equal(t, int64(2038074751), sieve.NthPrime(100000000)) //not required, just a fun challenge
}

func BenchmarkNewSieve(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sieve := NewSieve()
		assert.Equal(b, int64(2038074751), sieve.NthPrime(100000000))
	}
}

func FuzzNthPrime(f *testing.F) {
	sieve := NewSieve()

	f.Fuzz(func(t *testing.T, n int64) {
		fmt.Println(n)
		//The fuzzer is passing in negative values,
		//   If we were going to production with this code I'd want to
		//   get the interface changed and return an error if this happens
		n = int64(math.Abs(float64(n)))
		if !big.NewInt(sieve.NthPrime(n)).ProbablyPrime(0) {
			t.Errorf("the sieve produced a non-prime number at index %d", n)
		}
	})
}
