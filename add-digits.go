func addDigits(num int) int {
    if num < 10 {
        return num
    }
    if ans := num % 9;ans==0{
        return 9
    } else {
        return ans
    }
}
