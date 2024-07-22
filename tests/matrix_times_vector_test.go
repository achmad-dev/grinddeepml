package tests_test

import (
	"reflect"
	"testing"

	"github.com/achmad-dev/grinddeepml/problems"
)

// Using Generic type for test case struct
type matrixDotVectorTestCase[T int | float64] struct {
	name     string
	A        [][]T
	B        []T
	Expected []T
	WantErr  bool
}

// test matrixdotvector for int type
func TestMatrixDotVectorInt(t *testing.T) {
	testCases := []matrixDotVectorTestCase[int]{
		{
			name:     "valid 2x2 matrix with 2x1 vector",
			A:        [][]int{{1, 2}, {2, 4}},
			B:        []int{1, 2},
			Expected: []int{5, 10},
			WantErr:  false,
		},
		{
			name:     "invalid: vector length mismatch",
			A:        [][]int{{1, 2, 3}, {4, 5, 6}},
			B:        []int{1, 2},
			Expected: nil,
			WantErr:  true,
		},
		{
			name:     "empty matrix",
			A:        [][]int{},
			B:        []int{1, 2, 3},
			Expected: nil,
			WantErr:  true,
		},
	}

	runTests(t, testCases)
}

// test matrixdotvector for float type
func TestMatrixDotVectorFloat(t *testing.T) {
	testCases := []matrixDotVectorTestCase[float64]{
		{
			name:     "Invalid: vector length mismatch",
			A:        [][]float64{{1.1, 2.2, 3.3}, {4.4, 5.5, 6.6}},
			B:        []float64{1.0, 2.0},
			Expected: nil,
			WantErr:  true,
		},
		{
			name:     "valid 2x2 matrix with 2x1 vector",
			A:        [][]float64{{1.0, 2.0}, {2.0, 4.0}},
			B:        []float64{1.0, 2.0},
			Expected: []float64{5.0, 10.0},
			WantErr:  false,
		},
	}

	runTests(t, testCases)
}

// runTests is a helper function to run tests for both int and float64 types
func runTests[T int | float64](t *testing.T, testCases []matrixDotVectorTestCase[T]) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := problems.MatrixDotVector(tc.A, tc.B)

			if tc.WantErr {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(result, tc.Expected) {
					t.Errorf("Expected %v, but got %v", tc.Expected, result)
				}
			}
		})
	}
}
