package stringutils

// ContainsInSlice serve caller to check target in string slice
func ContainsInSlice(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}

	return false
}
