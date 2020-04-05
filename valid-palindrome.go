import "strings"
func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    for i,j:=0,len(s)-1;i<=j;{
        il := isAlphaNum(s[i])
        jl := isAlphaNum(s[j])
        if il && jl {
            if s[i]!=s[j] {
            return false
            } else {
                i++
                j--
            }
        }
        if !il{
            i++
        }
        if !jl{
            j--
        }
    }
    return true
}

func isAlphaNum(ch byte) bool {
    if ch >='a' && ch <='z' || ch>='0' && ch<='9'{
        return true
    }
    return false
}
