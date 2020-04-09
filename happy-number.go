func isHappy(n int) bool {
    fast := next(n)
    for n!=1 && fast!=1{
        if n == fast {
            return false
        }
        n = next(n)
        fast = next(next(fast))
        
    }
    return true
}

func next(n int) int {
    sum:=0 
    for n>0 {
        digit:=n%10
        sum+=digit*digit
        n/=10
    }
    return sum
}


