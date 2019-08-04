package sieve

// Sieve finds all the primes from 2 up to a given number using Sieve of Eratosthenes.
func Sieve(limit int) (primes []int) {
	if limit <= 1 {
		return
	}

	marker := make([]bool, limit+1)

	for p := 2; p*p <= limit; p++ {
		if !marker[p] {
			for i := p * p; i <= limit; i += p {
				marker[i] = true
			}
		}
	}

	for i := 2; i <= limit; i++ {
		if !marker[i] {
			primes = append(primes, i)
		}
	}

	return
}
