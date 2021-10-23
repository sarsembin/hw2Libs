package hw2Libs

import "math"

func RegisterFlip(s string) string{
	newS := ""
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			newS += string(s[i] + 32)
		} else if s[i] >= 'a' && s[i] <= 'z' {
			newS += string(s[i] - 32)
		}
	}
	return newS
}

func SquareRoot(num float64) float64 {
	return math.Sqrt(num)
}
