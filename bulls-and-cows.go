import "fmt"
func getHint(secret string, guess string) string {
    A:=0
    B:=0
    hash := make(map[byte]int,len(secret))
    for i,ll:=0,len(secret); i<ll; i++{
        hash[secret[i]]++
    }
    for i:=0; i<len(guess);{
        if guess[i]==secret[i] {
            A++
            hash[secret[i]]--
            if hash[secret[i]] == 0{
                delete(hash, secret[i])
            }
            guess = guess[:i]+guess[i+1:]
            secret = secret[:i]+secret[i+1:]
        } else {
            i++
        }
    }
    for i,ll:=0,len(guess); i<ll; i++{
        if _, ok := hash[guess[i]]; ok{
            B++
            hash[guess[i]]--        
            if hash[guess[i]]==0 {            
                delete(hash, guess[i])
            }
            
        }
    }
    return fmt.Sprintf("%dA%dB",A,B)
}
