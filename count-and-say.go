import "fmt"
func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    var ans string
    var last rune
    var count = 0 
    for _,r := range countAndSay(n-1) {
        if r == last {
            count++
        } else {
            if count != 0 {
            ans+=fmt.Sprintf("%d%c",count,last)
            }
            last = r
            count = 1
        }
    }
    if count != 0 {
    ans+=fmt.Sprintf("%d%c",count,last)
    }
    return ans
}
