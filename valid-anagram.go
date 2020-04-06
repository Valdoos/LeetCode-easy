func isAnagram(s string, t string) bool {
    if len(s)!=len(t) {
        return false
    }
    hashS := make(map[byte]int)
    for i,end:=0,len(s);i<end;i++{
        hashS[s[i]]++
    }
    for i,end:=0,len(t);i<end;i++{
        hashS[t[i]]--
    }
    for _,v := range hashS{
        if v != 0 {
            return false
        }
    }
    return true
}
