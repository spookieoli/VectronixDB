package utils

// IsIn checks if wrd is in arr
func IsIn(wrd string, arr []string) bool {
	for _, v := range arr {
		if v == wrd {
			return true
		}
	}
	return false
}
