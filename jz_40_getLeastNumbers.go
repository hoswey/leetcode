func getLeastNumbers(arr []int, k int) []int {
    
    if k >= len(arr) {
        return arr
    }
    
    if len(arr) == 0 || k == 0 {
        return nil
    }
    
    start := 0
    end := len(arr) - 1
    
    for {        
        
        p := partitions(arr, start, end)  
        if p == k - 1 {
            return arr[0: k]
        }
        
        if p < k - 1 {
            start = p + 1
        } else {
            end = p - 1
        }
    }

}

func partitions(arr []int, lo, hi int) int {
    
    if lo == hi {
        return lo
    }
    
    mid := lo + (hi - lo) >> 1
    
    if arr[lo] > arr[mid] {
        arr[lo], arr[mid] = arr[mid], arr[lo]
    }
    
    if arr[hi] < arr[lo] {
        arr[hi], arr[lo] = arr[lo], arr[hi]
    }
    
    if arr[mid] < arr[hi] {
        arr[mid], arr[hi] = arr[hi], arr[mid]
    }
    
    pivot := arr[hi]
    
    i := lo
    
    for j := lo; j <= hi - 1; j++ {
        
        if arr[j] <= pivot {
            arr[i], arr[j] = arr[j], arr[i]
            i++
        }                
    }
    arr[i], arr[hi] = arr[hi], arr[i]
    return i
}