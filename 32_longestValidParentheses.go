import "container/list"


var nilVal = -10

type Stack struct {
	l *list.List
}

func NewStack() *Stack {

	s := Stack{l: list.New()}
	return &s
}

func (s *Stack) Peek() int {

	if s.l.Len() == 0 {
		return nilVal
	}

	return s.l.Front().Value.(int)
} 

func (s *Stack) Pop() int {

	if s.l.Len() == 0 {
		return nilVal
	}

	ele := s.l.Front()
	s.l.Remove(ele)
	return ele.Value.(int)
} 

func (s *Stack) Empty() bool {
	return s.l.Len() == 0
}


func (s *Stack) Push(v int) {
	s.l.PushFront(v)
}



func longestValidParentheses(s string) int {

	stack := NewStack()
	stack.Push(-1)

	var ans int
	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			stack.Push(i)
		} else {
			_ = stack.Pop()
			if stack.Empty() {
				stack.Push(i)
			} else {
				ans = max(ans, i - stack.Peek())
			}

		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}