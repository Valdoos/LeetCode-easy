import "fmt"
func addStrings(num1 string, num2 string) string {

    n1 := &num1
    n2 := &num2
    if len(num1) < len(num2) {
        n1,n2 = n2,n1
    }
    l1 := len(*n1)
    l2 := len(*n2)
    l := max(l1,l2)+1
    ans := make([]byte,l)
    var supple byte
    l1,l2,l = l1-1,l2-1,l-1
    for l2 >= 0 {
        ans[l]=(*n1)[l1]+ (*n2)[l2] - '0' + supple
        if ans[l] > '9' {
            ans[l]-=10
            supple = 1
        } else {
            supple = 0
        }
        l1--
        l2--
        l--
    }
    for l1 >= 0 {
        ans[l]=(*n1)[l1] + supple
        if ans[l] > '9' {
            ans[l]-=10
            supple = 1
        } else {
            supple = 0
        }
        l1--
        l--
    }
    if supple == 0 {
        return string(ans[1:])
    }
    ans[0] = '0'+ supple
    return string(ans)
    
} 
