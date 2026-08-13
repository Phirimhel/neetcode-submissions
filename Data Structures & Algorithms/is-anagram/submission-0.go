func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false 
	}
	
	set := make(map[rune]int, len(s))
	for _, r := range s {
		set[r]++
	}

	for _, r := range t {
		if _ , exists := set[r]; !exists {
			return false
		}
		set[r]--
		if set[r] == 0 {
			delete(set, r)
		}
	}
	return true
}
