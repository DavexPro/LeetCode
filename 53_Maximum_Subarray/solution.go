func maxSubArray(nums []int) int {
    
    if nums == nil {
        return 0
    }
    
    max := nums[0]
    curr := 0
    
    for _, v := range nums {
        curr = Max(v, curr + v)
        max = Max(max, curr)
    }
    
    return max
}

func Max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
