func findDuplicate(nums []int) int {
    
    if nums == nil || len(nums) < 2 {
        return -1
    }
    
    for i, p := range nums {
        for _, q := range nums[i+1:] {
            if p == q {
                return q
            }
        }
    }
    
    return -1
}
