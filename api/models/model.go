package models

func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func Contains(slices []string, text string) bool {
	for _, value := range slices {
		if text == value {
			return true
		}
	}
	return false
}
