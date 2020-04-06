func convertToTitle(n int) string {
    if n == 0 {
        return ""
    }
    if n<=26 {
        return string('A'+n-1)
    }
    return convertToTitle((n-1)/26)+string('A'+(n-1)%26)
}
