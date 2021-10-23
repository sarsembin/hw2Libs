package hw2Libs

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
