func toHex(num int) string {
    if num == 0 {
        return "0"
    }
    var n uint64
    if num < 0 {
        n = 1<<32 + uint64(num)
    } else {
        n = uint64(num)
    }
    hex := "0123456789abcdef"
    var ans string 
    for n > 0 {
        x := n%16
        ans=string(hex[x])+ans
        n/=16
    }
    return ans
}
