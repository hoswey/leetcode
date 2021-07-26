/*
 * @lc app=leetcode id=241 lang=golang
 *
 * [241] Different Ways to Add Parentheses
 */

// @lc code=start
import "strconv"

func diffWaysToCompute(expression string) []int {
    
	var ways []int

	for i := 0; i < len(expression); i++ {
		c := expression[i]
		if c == '+' || c == '-' || c == '*' {
			lWays := diffWaysToCompute(expression[0:i])
			rWays := diffWaysToCompute(expression[i+1:])
			for _, lWay := range lWays {
				for _, rWay := range rWays {
					switch c {
					case '+': ways = append(ways, lWay + rWay);
					case '-': ways = append(ways, lWay - rWay);
					case '*': ways = append(ways, lWay * rWay);
					}
				}
			}
		}
	}

	if ways == nil {
		num, _ := strconv.Atoi(expression)
		ways = append(ways, num)
	}
	return ways
}
// @lc code=end

