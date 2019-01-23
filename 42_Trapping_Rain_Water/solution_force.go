func trap(height []int) int {
    if height == nil || len(height) == 0 {
        return 0
    }
    
    ret := 0
    for i:=1;i<len(height)-1;i++ {
        max_left := 0
        max_right := 0
        // left side
        for j:=i;j>=0;j-- {
            max_left = max(max_left, height[j])
        }
        // right side
        for j:=i;j<len(height);j++ {
            max_right = max(max_right, height[j])
        }
        ret += min(max_left, max_right) - height[i]
    }
    return ret
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a int, b int) int {
    if a > b {
        return b
    }
    return a
}
