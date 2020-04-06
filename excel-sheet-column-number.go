func titleToNumber(s string) int {
    l :=len(s)
    if l==1{
        return int(s[0]-'A')+1
    }
    return 26*titleToNumber(s[:l-1])+(int(s[l-1]-'A')+1)
}
