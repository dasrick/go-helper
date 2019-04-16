package gohelper

// StringDeduplicate ... deduplicate a slice of strings
func StringDeduplicate(s []string) []string {
	if len(s) == 0 {
		return s
	}

	result := make([]string, 0)
	seen := make(map[string]string)
	for _, val := range s {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}
