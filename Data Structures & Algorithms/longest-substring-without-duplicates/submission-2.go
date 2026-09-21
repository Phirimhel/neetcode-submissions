func lengthOfLongestSubstring(s string) int {

	characters := make(map[byte]int)
	maxLen := 0 

	l, r := 0, 0 

	
	for r < len(s) {

		if lastIndex, ok := characters[s[r]]; ok && lastIndex >= l {
			l = lastIndex + 1
		} 

		characters[s[r]] = r

		currentLen := r - l + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}

		r++
	}

	return maxLen
}
