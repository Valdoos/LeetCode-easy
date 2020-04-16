import "fmt"
func fizzBuzz(n int) []string {
    ans  := make([]string,n)
    nextThree := 3
    nextFive := 5
    for i:=1;i<=n;i++ {
        out := ""
        if i == nextThree {
            out+="Fizz"
            nextThree+=3
        }
        if i == nextFive {
            out+="Buzz"
            nextFive+=5
        }
        if len(out) == 0 {
            out=fmt.Sprintf("%d",i)
        } 
        ans[i-1]=out
    }
    return ans
}
