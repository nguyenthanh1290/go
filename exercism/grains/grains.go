// Package grains provides function that calculates the number of grains of wheat on a chessboard given that the number on each square doubles.
package grains

import "errors"

// Square returns the number of grains on a square.
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("square must be from 1 to 64")
	}

	return 1 << uint64(square-1), nil
}

// Total returns the total number of grains there could be on a chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
