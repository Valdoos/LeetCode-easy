func isIsomorphic(s string, t string) bool {
    l := len(s)
    if l != len(t) {
        return false
    }
    return halfIsomorphic(s,t,l) && halfIsomorphic(t,s,l)
}

func halfIsomorphic(s string, t string, l int) bool {
    hash := make([]byte, 1<<7, 1<<7)
    for i:=0; i<l; i++ {
        if hash[s[i]] != 0 {
            if  hash[s[i]] != t[i] {
                return false
            }
        }
        hash[s[i]] = t[i]
    }
    return true
}
