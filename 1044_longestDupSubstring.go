import "strings"

func longestDupSubstring(S string) string {

    var ans string

    low := 0
    high := len(S) - 1

    for low <= high {

        mid := low + (high - low)/2

        ret := search(S,mid)
        if ret == "" {
            high = mid - 1
        } else {
            low = mid + 1
            ans = ret
        }
    }
    return ans
}

func search(S string, length int) string {

    mod :=  1 << 31 - 1
    d := 256

    h := 1
    for i := 1; i < length; i++ {
        h =(h * d) % mod
    }

    seen := make(map[int]bool)

    var now int
    for i := 0; i < length; i++{
        now = (now * d % mod + int(S[i])) % mod
    }
    seen[now] = true

    for i := length; i < len(S); i++ {

        now = (now - int(S[i-length]) * h % mod + mod) % mod
        now = (now * d + int(S[i])) % mod

        if seen[now] && strings.Contains(S[:i], S[i - length + 1 : i + 1]) {
            return S[i - length + 1 : i + 1]
        }
        seen[now] = true
    }
    return ""
}