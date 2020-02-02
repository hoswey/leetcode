var memo map[string]int
var S []int
var totalPick, memoPick int

func initVar() {
	memo = make(map[string]int)

	//存储从i到结束的石头总数	
	S = make([]int, len(piles))
	totalPick, memoPick = 0, 0	
}

func stoneGameII(piles []int) int {

	l := len(piles)  

	S[l - 1] = piles[l - 1]
	for i := len(piles) - 2;  i >= 0; i-- {
		S[i] = piles[i] + S[i+1]
	}

	ans := pick(piles, 0, 1)
    fmt.Printf("piles:%v,ans:%d,totalPick:%d, memoPick:%d\n", piles, ans, totalPick, memoPick)
	return ans
}

func pick(piles []int, i, M int) int{

	k := fmt.Sprintf("%d:%d", i, M)

	totalPick++
	if val, ok := memo[k]; ok {
		memoPick++
		return val
	}

	l := len(piles)

	if i >= l {
		return 0
	}

	if M * 2 >= l - i {
		var sum int
		for _, val := range piles[i:] {
			sum += val
		}
		memo[k] = sum
		return sum
	}

	var best int 
	for x := 1; x <= M * 2; x++ {
		best = max(best, S[i] - pick(piles, i + x, max(x, M)))
	}
	memo[k] = best
    return best
}

func max(a, b int) int{

	if a > b {
		return a
	}
	return b
}