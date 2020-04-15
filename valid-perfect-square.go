func isPerfectSquare(num int) bool {
    low:=1
    n := num
    for low<num-1{
        middle := (num+low)/2
        middleSquare := middle*middle
        if middleSquare==n {
            return true
        } else if middleSquare<n{
            low = middle
        } else {
            num = middle
        }
    }
    return low*low==n
}
