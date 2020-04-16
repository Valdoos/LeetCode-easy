import "fmt"
func readBinaryWatch(num int) []string {
    var res []string
    doRead(&res, num, 0, 0, 0)
    return res
}

func doRead(res *[]string, num, hours, minutes, pos int) {
    if num == 0 {
        (*res) = append((*res),fmt.Sprintf("%d:%02d",hours,minutes))
        return
    }
    for i:=pos;i < 10;i++ {
        if i >= 4 {
            if minutes+1<<(i-4)>=60{
                return 
            }
            doRead(res,num-1,hours,minutes+1<<(i-4),i+1)
        } else {
            if hours+1<<i > 11 {
                continue
            }
            doRead(res,num-1,hours+1<<i,minutes,i+1)
        }
    }
}
