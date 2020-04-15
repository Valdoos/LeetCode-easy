/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guessNumber(n int) int {
    low:=1
    high:=n
    for low<high{
        medium := (low+high)>>1
        result := guess(medium)
        if result == 0 {
            return medium
        } else if result == -1 {
            high = medium-1
        } else {
            low = medium+1
        }
    }
    return low
}
