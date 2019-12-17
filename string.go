package gohelper

// StringDeduplicate ... deduplicate a slice of strings
func StringDeduplicate(ss []string) []string {
	result := make([]string, 0)
	seen := make(map[string]string)
	for i := range ss {
		if _, ok := seen[ss[i]]; !ok {
			result = append(result, ss[i])
			seen[ss[i]] = ss[i]
		}
	}
	return result
}
