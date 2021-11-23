/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
import "container/list"

func evalRPN(tokens []string) int {

	l := list.New()

	for _, token := range tokens {

		switch token {

		case "+":
			o1, o2 := getOperand(l)
			l.PushBack(o1 + o2)
		case "-":
			o1, o2 := getOperand(l)
			l.PushBack(o1 - o2)
		case "*":
			o1, o2 := getOperand(l)
			l.PushBack(o1 * o2)
		case "/":
			o1, o2 := getOperand(l)
			l.PushBack(o1 / o2)
		default:
			l.PushBack(strToInt(token))
		}
	}

	return l.Front().Value.(int)
}

func getOperand(l *list.List) (int, int) {

	e := l.Back()
	l.Remove(e)
	o2 := e.Value.(int)

	e = l.Back()
	l.Remove(e)
	o1 := e.Value.(int)

	return o1, o2
}

func strToInt(str string) int {

	var ret int

	var negative bool
	for i := 0; i < len(str); i++ {

		if str[i] == '-' {
			negative = true
		} else {
			ret = ret*10 + int(str[i] - '0')
		}
	}

	if negative {
		ret = - ret
	}
	return ret
}
// @lc code=end

