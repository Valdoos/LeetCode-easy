func longestPalindrome(s string) string {
    if len(s)<= 1 {
        return s
    }
    start, end := 0, 0
    for i,l:=0,len(s); i<l; i++{
        l1 := expand(s,i,i)
        l2 := expand(s,i,i+1)
        l3 :=l1
        if l3 < l2 {
            l3 = l2
        }
        if l3 > end-start{
            start = i - (l3-1)/2
            end = i + l3/2
        }
    }
    return s[start:end+1]
}

func expand(s string, left, right int) int {
    l,r:= left,right
    for  l>=0 && r<len(s) && s[l]==s[r] {
        l--
        r++
    }
    return r-l-1
}
