package main

import "fmt"
import "container/list"


type Stack struct {
	l *list.List	
}

func NewStack() *Stack {

	return &Stack{l: list.New()}
}

func (s *Stack) pop() int {
	ele := s.l.Front()
	s.l.Remove(ele)
	return ele.Value.(int)
}

func (s *Stack) peek() int {
	ele := s.l.Front()
	return ele.Value.(int)	
}

func (s *Stack) len() int {
	return s.l.Len()
}

func (s *Stack) push(v int) {
	s.l.PushFront(v)
}

func (s *Stack) String() string {

	ele := s.l.Front()
	if ele == nil {
		return ""
	}

	var slice []int
	for ele != nil {
		slice = append(slice, ele.Value.(int))
		ele = ele.Next()
	}

	return fmt.Sprintf("%v", slice)
}


func largestRectangleArea(heights []int) int {

	if len(heights) == 0 {
		return 0
	}

	s := NewStack()
	s.push(-1)

	var ans int
	for i, _ := range heights {
		
		for s.len() > 1 && heights[s.peek()] > heights[i] {

			fmt.Printf("i:%d, stack:%v\n",i, s)
			height := heights[s.pop()]
			area := height * (i - s.peek() - 1) 

			fmt.Printf("i:%d, height:%v, area:%v\n", i,height, area)
			if area > ans {
				ans = area
			}
		}
		s.push(i)
	}

	for s.len() > 1 {
		height := heights[s.pop()]
		area := height * (len(heights) - s.peek() - 1) 
		if area > ans {
			ans = area
		}
	}

	return ans
}

func main() {


	var input []int
	input = []int{2,1,5,6,2,3}

	fmt.Printf("input:%v, ouput:%v, expected:%d", input, largestRectangleArea(input), 10)

	input = []int{2}

	fmt.Printf("input:%v, ouput:%v, expected:%d", input, largestRectangleArea(input), 2)

	input = []int{2,1,1,1,2,3}

	fmt.Printf("input:%v, ouput:%v, expected:%d", input, largestRectangleArea(input), 6)	

}