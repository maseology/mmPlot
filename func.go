package mmplt

import "math"

// rev is quick function used to reverse order of a slice
func rev(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// revF is quick function used to reverse order of a float64 slice
func revF(s []float64) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// onlyPositive removes all value <= 0.0 and all NaN's
func onlyPositive(s []float64) []float64 {
	var x []int
	for i := range s {
		if s[i] <= 0 || math.IsNaN(s[i]) {
			x = append(x, i)
		}
	}
	rev(x)
	for _, i := range x {
		s = append(s[:i], s[i+1:]...)
	}
	return s
}
