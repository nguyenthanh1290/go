package perfect

import (
	"errors"
)

// Classification represents the classification values for natural numbers.
type Classification int

// Classification values
const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

// ErrOnlyPositive is a customer error type.
var ErrOnlyPositive = errors.New("positive integer is required")

// Classify determines if a number is perfect, abundant, or deficient based on
// Nicomachus' (60 - 120 CE) classification scheme for natural numbers.
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return -1, ErrOnlyPositive
	}

	aSum := aliquotSum(n)
	if aSum == n {
		return ClassificationPerfect, nil
	}
	if aSum > n {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}

// The aliquot sum is the sum of the factors of a number not including the number itself.
func aliquotSum(n int64) (aSum int64) {
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			aSum += i
		}
	}
	return
}
