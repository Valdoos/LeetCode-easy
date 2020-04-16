func countSegments(s string) int {
    inSegment := false
    count:=0
    for _,ch := range []byte(s) {
        if ch == ' ' {
            if inSegment {
                count++
                inSegment = false
            }
        } else {
            inSegment = true
        }
    }
    if inSegment {
        count++
    }
    return count
}
