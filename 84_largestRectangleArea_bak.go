package main

import "fmt"
import "container/list"

var EmptyStack = 99

type Stack struct {
    l *list.List
}

func (s *Stack) Pop () int {

    if s.l.Len() == 0 {
        return EmptyStack
    }

    ele := s.l.Front()
    s.l.Remove(ele)

    return ele.Value.(int)
}

func (s *Stack) Push(val int) {
    s.l.PushFront(val)
}

func (s *Stack) Peek () int {

    if s.l.Len() == 0 {
        return EmptyStack
    }

    return s.l.Front().Value.(int)
}


func (s *Stack) Len() int {

    return s.l.Len()
}


func largestRectangleArea(heights []int) int {
    
    var ans int

    l :=  Stack{l: list.New()}
    l.Push(-1)

    for i := 0; i < len(heights); i++ {

        for l.Len() > 1 && heights[l.Peek()] > heights[i] {
            ans = max(ans, heights[l.Pop()]  * (i - 1 - l.Peek()))
        }
        l.Push(i)
    }

    for l.Len() > 1 {
        ans = max(ans, heights[l.Pop()] * (len(heights) - 1  - l.Peek()))
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

func main() {

    input := []int{2,1,2}
    fmt.Printf("input:%v, output:%v\n", input, largestRectangleArea(input))
}