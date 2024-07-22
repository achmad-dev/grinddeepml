package problems

import (
	"errors"
	"fmt"
)

func MatrixDotVector[T int | float64](a [][]T, b []T) ([]T, error) {
	if len(a) == 0 || len(a[0]) == 0 || len(b) == 0 {
		return nil, errors.New("matrix empty")
	}
	if len(a[0]) != len(b) {
		return nil, fmt.Errorf("length of a(%d) and b(%d) is not same", len(a[0]), len(b))
	}
	res := make([]T, len(a))
	for i := range a {
		//hold := 0
		for j := range b {
			res[i] += a[i][j] * b[j]
		}
	}
	return res, nil
}
