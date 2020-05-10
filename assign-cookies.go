import "sort"
func findContentChildren(g []int, s []int) int {
    ans := 0
    sort.Ints(g)
    sort.Ints(s)
    for i,j := 0,0;i<len(g) && j < len(s); {
        if g[i]<=s[j] {
            i++
            j++
            ans++
        } else {
            j++
        }
    }
    return ans
}
