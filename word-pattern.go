import "strings"
func wordPattern(pattern string, str string) bool {
    words := strings.Split(str," ")
    if len(words)!=len(pattern) {
        return false
    }
    hash := make([]string,1<<7,1<<7)
    used := make(map[string]struct{},len(words))
    for i, s := range words {
        if hash[pattern[i]] == ""{
            if _,ok := used[s]; ok{
                return false
            }
            used[s] = struct{}{}
            hash[pattern[i]] = s
        } else if hash[pattern[i]] != s {
            return false
        }
    }
    
    return true
}
