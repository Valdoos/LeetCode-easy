import "fmt"
func addStrings(num1 string, num2 string) string {
    const max = 5100
    str1 := make([]byte,max-len(num1),max)
    str1 = []byte(string(str1)+num1)
    str2 := make([]byte,max-len(num2),max)
    str2 = []byte(string(str2)+num2)
    ans := make([]byte,max)
    l := helper(str1,str2,ans,max-1,maxInt(len(num1),len(num2)),0)
    return string(ans[l+1:])
}
func maxInt(a, b int) int {
    if a > b {
        return a
    } 
    return b
} 
func helper(num1 , num2 []byte,ans []byte, l , end int, add byte) int { 
    a := num1[l]
    b := num2[l]
    sum := a+b+add
    if a!=0{
        sum-='0'
    }
    if b!=0{
        sum-='0'
    }
    ans[l] = sum%10+'0'
    if end == -1 || (end==0 && sum%10==0){
        return l
    }
    return helper(num1, num2, ans, l-1,end-1, sum/10)
}
