func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true 
	}
	
	s = strings.ToLower(s)
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = reg.ReplaceAllString(s, "")

	
	

	left := 0 
	right := len(s) - 1 
	for left <= right {

		if s[left] != s[right] {
			return false
		}

		left++ 
		right-- 

	}
	return true
}
