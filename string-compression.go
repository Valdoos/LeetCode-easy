import "math"
import "fmt"
func compress(chars []byte) int {
    if len(chars) ==  0{
        return 0
    }
    ind := 0
    count := 0
    last := chars[0]
    for _,ch := range chars {
        if ch != last {
            chars[ind]=last
            if count > 1 {
                k:=1
                for count>=int(math.Pow(10,float64(k))){
                    k++
                }
                ind+=k
               
                for count > 0 {
                    chars[ind] = '0' + byte(count%10)

                    count/=10
                    ind--
                }
                ind+=k+1
            } else {
                ind++
            }
            count = 1
            last = ch
        } else {
            count++
        }
        
    }
      chars[ind]=last
            if count > 1 {
                k:=1
                for count>=int(math.Pow(10,float64(k))){
                    k++
                }
                ind+=k
               
                for count > 0 {
                    chars[ind] = '0' + byte(count%10)

                    count/=10
                    ind--
                }
                ind+=k+1
            } else {
                ind++
            }
    return ind
}
