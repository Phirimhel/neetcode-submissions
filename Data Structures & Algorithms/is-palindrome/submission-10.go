func isPalindrome(s string) bool {
    n := len(s)

    s = strings.ToLower(s)
    left := 0
    right := n-1 
    for left <= right { 
        if !isAlphanumeric(s[left]) {
            left++
            continue
        }

        if !isAlphanumeric(s[right]) {
            right--
            continue
        }

        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}
func isAlphanumeric(b byte) bool {
    return (b >= 'a' && b <='z') || (b >= '0' && b <='9')
}