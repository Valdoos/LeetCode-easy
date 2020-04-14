/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    first := 0
    for first < n-1{ 
        medium := (n+first)/2
        if isBadVersion(medium) {
            n = medium
        } else {
            first = medium
        }
    }
    return n
}
