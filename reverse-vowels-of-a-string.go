func reverseVowels(s string) string {
    str := []byte(s)
    isVowel := [128]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true, 'A':true, 'E':true, 'I':true, 'O':true, 'U':true}
    for i,j:=0,len(str)-1;i<j;{
        if !isVowel[str[i]] {
            i++
            continue
        }
        if !isVowel[str[j]] {
            j--
            continue
        }
        str[i], str[j] = str[j], str[i]
        i++
        j--
    }
    return string(str)
}
