func addBinary(a string, b string) string {
    b1 := []byte(a)
    b2 := []byte(b)
    ans := make([]byte,max(a,b)+1)
    if len(a)<len(b){
        b1,b2=b2,b1
    }
    for i:=range ans {
        ans[i]='0'
    }
    i,j,k:=len(b1)-1,len(b2)-1,len(ans)-1
    for ;i>=0 && j>=0;i,j,k=i-1,j-1,k-1 {
        ans[k]+=b1[i]+b2[j]-'0'-'0'
        if ans[k]>='2' {
            ans[k-1]++
            ans[k]=ans[k]-2
        }
    }
    for ;i>=0;i,k=i-1,k-1 {
         ans[k]+=b1[i]-'0'
         if ans[k]>='2' {
            ans[k-1]++
            ans[k]=ans[k]-2
         }
    }
    if ans[0]=='0'{
        return string(ans[1:])
    }
    return string(ans)
}

func max(s1,s2 string) int {
    if len(s1)>len(s2) {
        return len(s1)
    }
    return len(s2)
}
