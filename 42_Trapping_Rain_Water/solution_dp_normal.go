func trap(height []int) int {
    if height == nil || len(height) == 0 {
        return 0
    }

    ret := 0
    leftMax := make([]int, len(height))
    rightMax := make([]int, len(height))
    
    // find the MAX height from the left
    leftMax[0] = height[0]
    for i:=1;i<len(height);i++ {
        leftMax[i] = max(height[i], leftMax[i-1])
    }
    
    // find the MAX height from the right
    rightMax[len(height)-1] = height[len(height)-1]
    for i:=len(height)-2;i>=0;i-- {
        rightMax[i] = max(height[i], rightMax[i+1])
    }
    
    // combine the ret, and based on the less one
    for i:=1;i<len(height)-1;i++ {
        ret += min(leftMax[i], rightMax[i]) - height[i]
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
