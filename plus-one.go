func plusOne(digits []int) []int {
    i := len(digits)-1
    for ;i>=0;i--{
        digits[i]=(digits[i]+1)%10
        if digits[i]!=0{
            break
        }
    }
    if digits[0]==0 && i==-1 {
        return append([]int{1},digits...)
    } 
    return digits
}
