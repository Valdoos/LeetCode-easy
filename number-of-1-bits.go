func hammingWeight(num uint32) int {
    if num == 0 {
        return 0
    }
    if num&1==1{
        return 1 + hammingWeight(num>>1) 
    }
    return hammingWeight(num>>1)
}
