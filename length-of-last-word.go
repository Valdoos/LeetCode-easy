func lengthOfLastWord(s string) int {
    if s == ""{
        return 0
    }
    var i =len(s)-1
    var end int
    for i>=0 && s[i]==' '{
        i--
    }
    end = i+1
    i--
    for i>=0{
        if s[i] == ' ' {
            return end-i-1
        }
        i--
    }
    return end
    
}
