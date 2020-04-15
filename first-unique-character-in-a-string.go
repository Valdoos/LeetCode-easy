func firstUniqChar(s string) int {
    var hash [32]int
    for _,ch := range []byte(s) {
        hash[ch-'a']++
    }
     for i,ch := range []byte(s) {
         if hash[ch-'a'] == 1 {
             return i
         }
    }
    return -1
}
