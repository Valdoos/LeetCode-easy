func isSubsequence(s string, t string) bool {
    if len(t)<len(s) {
        return false
    }
    i:=0
    for _,ch := range []byte(t) {
        if i == len(s) {
            return true
        } 
        if ch == s[i] {
            i++
        }
    }
    return i == len(s)
}
