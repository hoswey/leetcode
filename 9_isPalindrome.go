import "strconv"

func isPalindrome(x int) bool {

    s := strconv.Itoa(x)

    l := len(s)
    for i := 0; i < l/2; i++ {
        if s[i] != s[l - i - 1] {
            return false
        }
    }
    return true
}
