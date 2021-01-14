func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    ans := 0
    hash := make([]int,128)
    i:=0
    for j,value := range s {
        i = max(i, hash[value])
        ans = max(ans,j-i+1)
        hash[value]=j+1
    }
    return ans
}
func max(a,b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
