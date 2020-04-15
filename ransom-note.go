func canConstruct(ransomNote string, magazine string) bool {
    if len(ransomNote)>len(magazine) {
        return false
    }
    var hash [128]byte
    count := 0
    for _, ch:=range []byte(ransomNote) {
        hash[ch]++
        count++
    }
    for _, ch:=range []byte(magazine) {
        if count == 0 {
            return true
        }
        if hash[ch] > 0 {
            hash[ch]--
            count--
        }
    }
    return count==0
}
