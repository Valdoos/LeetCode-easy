func longestPalindrome(s string) int {
    ans := 0
    medium :=0
    var hash [128]bool
    for _, ch := range []byte(s) {
        hash[ch]=!hash[ch]
        if hash[ch]==false{
            ans+=2
            medium--
        } else {
            medium++
        }
    }
    if medium > 0 {
        return ans+1
    }
    return ans
}
