//给出一个区间的集合，请合并所有重叠的区间。

//示例 1:

//输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2:

//输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

package main

import (
  "sort"
  "container/list"
)

import (
  "sort"
  "container/list"
)

func merge(intervals [][]int) [][]int {

  if len(intervals) <= 1 {
    return intervals
  }

  sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
  })

  l := list.New()
  l.PushBack(intervals[0])

  for i := 1; i < len(intervals); i++{

    last := l.Back().Value.([]int)

    if last[1] < intervals[i][0] {
      l.PushBack(intervals[i])
    } else {
      max := last[1]
      if max < intervals[i][1] {
        max = intervals[i][1]
      }
      last[1] = max
    }
  }

  var ans [][]int
  for l.Len() > 0 {
    ele := l.Front()
    l.Remove(ele)
    ans = append(ans, ele.Value.([]int))
  }
  return ans
}
