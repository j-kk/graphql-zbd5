package graph

func IsRGB(color *string) bool {
	s := *color
	if len(s) != 7 {
		return false
	}
	if s[0] != '#' {
		return false
	}
	for i := 1; i <= 6; i++ {
		if !(('0' <= s[i] && s[i] <= '9') || ('A' <= s[i] && s[i] <= 'F')) {
			return false
		}
	}
	return true
}
