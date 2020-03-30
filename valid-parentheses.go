func isValid(s string) bool {
    var stack []rune
    for _,b:= range s {
        if b == '(' || b=='[' || b =='{' {
            stack = append(stack,b)
        } else {
            n:=len(stack)-1
            if n < 0 {
                return false
            }
            diff:=b-stack[n]
            if diff > 2 || diff < -2 {
                return false
            }
            stack = stack[:n]
        }
    }
    if len(stack) > 0 {
        return false
    }
    return true
}
