import "strings"
func longestCommonPrefix(strs []string) string {
    if len(strs)==0 {
        return ""
    }
    prefix := strs[0]
    for i:=len(strs)-1;i>=1;i--{
        for !strings.HasPrefix(strs[i],prefix) {
            if prefix == "" {
            return ""
            }
            prefix=prefix[:len(prefix)-1]
        }

    }
    return prefix
}
