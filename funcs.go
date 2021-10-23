package hw2Libs

import (
	"math"
	"github.com/twharmon/gouid"
)

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

func GenerateUUID(length int) string{
	return gouid.String(length, gouid.MixedCaseAlpha)
}
