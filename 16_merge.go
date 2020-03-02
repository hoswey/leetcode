import "sort"

type Intervals struct {
    d [][]int
}

func (it *Intervals) Len() int{
    return len(it.d)
}

func (it *Intervals) Swap(i, j int)  {
    
    it.d[i], it.d[j] = it.d[j], it.d[i]
}

func (it *Intervals) Less(i, j int) bool {
    return it.d[i][0] < it.d[j][0]
}

func merge(intervals [][]int) [][]int {
    
    if len(intervals) <= 1 {
        return intervals
    }
    
    it := &Intervals{d: intervals}
    sort.Sort(it)
     //所有区间按 start排序
    //定义merge数组， 第一个区间进组
    //对于后面的每个区间， 获取merge数组最后的一个区间， 假如当前区间的start < 最后区间的end, 那么数组可以合并

    var merge [][]int
    merge = append(merge, intervals[0])
    
    for i := 1; i < len(intervals); i++ {
        
        cur := intervals[i]
        last := merge[len(merge)-1]
        
        if cur[0] > last[1] {
            merge = append(merge, cur)
        } else {
            
            if cur[1] <= last[1] {
                //当前的end小于上一个空间的end, 表示被包含   
            } else {
                last[1] = cur[1]
            }         
        }
    }
    return merge
}