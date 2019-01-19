func findDuplicate(nums []int) int {
    
    if nums == nil || len(nums) < 2 {
        return -1
    }
 
    start := 1
    end := len(nums) - 1
    
    for start < end {
        
        count := 0
        mid := (end - start)/2 + start
        
        for _, v := range nums {
            if v >= start && v <= mid {
                count ++
            }
        }
        
        if count <= (mid - start) + 1 {
            start = mid + 1
        } else {
            end = mid
        }
    }
    
    return end
}
