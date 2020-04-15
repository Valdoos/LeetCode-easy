import "fmt"
func getHint(secret string, guess string) string {
    A:=0
    B:=0
    var countS [10]int
    var countG [10]int
    for i := range secret {
        s := secret[i]-'0'
        g := guess[i]-'0'
        if secret[i] == guess[i] {
            A++
            continue
        }
        if countS[g] > 0 {
            B++
            countS[g]--
        } else {
            countG[g]++
        }
        if countG[s] > 0 {
            B++
            countG[s]--
        } else {
            countS[s]++
        }
    }
    return fmt.Sprintf("%dA%dB",A,B)
}
