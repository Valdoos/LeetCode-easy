func myAtoi(str string) int {
    var num uint
    var i int
    l := len(str)
    sign := true
    for i < l && str[i] == ' ' {
        i++
    }
    if i < l && str[i] == '-' {
        sign = false
        i++
    } else if i < l && str[i] == '+' {
        i++
    }
    for i < l && str[i] == '0' {
        i++
    }
    count := 0 
    for i < l && str[i] >= '0' && str[i] <= '9' {
        i++
        count++
    }
    i--
    var mult uint= 1
    if count > 10 {
         if sign {
                return 1 << 31 -1
        } else {
                return -1 << 31
        }
    }
    for count > 0 {
        num+=uint(str[i]-'0')*mult
        if num >= 1 << 31 - 1{
            if sign {
                return 1 << 31 -1
            } else if num >= 1 << 31 {
                return -1 << 31
            }
        }
        mult*=10
        count--
        i--
    }
    if sign {
        return int(num)
    } else {
        return int(-num)
    }
}
