func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    if len(needle) > len(haystack) {
        return -1
    }
    var j = 0 
    for i:=0;i<len(haystack);i++ {
        if haystack[i] == needle[j]{
            j++
            if j==len(needle) {
               return i-j+1
            }
        } else {
            i-=j
            j = 0
        }
    }
    return -1
}
