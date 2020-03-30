func romanToInt(s string) int {
    myMap := map[byte]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }
    previous := myMap[s[len(s)-1]]
    ans := previous
    for i:=len(s)-2;i>=0;i-- {
        currentSymbol := s[i]
        current := myMap[currentSymbol] 
        if current < previous {
            ans-=current
        } else {
            ans+=current
        }
        previous = current
    }
    return ans
}
