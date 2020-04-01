func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    start :=1
    end:= x/2
    for start<end-1{
        target := (start+end)/2 
        if target*target<x{
            start = target
        } else {
            end = target
        }
    }
    if end*end<=x {
        return end
    }
    return start
}
