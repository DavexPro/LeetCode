func maximalSquare(matrix [][]byte) int {
    if matrix == nil || len(matrix) == 0 {
        return 0
    }
    
    side := 0
    m, n := len(matrix[0]), len(matrix)
    
    // make dp matrix
    dp := make([][]int, n)
    for i :=0; i < n; i++ {
        dp[i] = make([]int, m)
    }

    // count the side length
    for x:=0;x<n;x++ {
        for y:=0;y<m;y++ {
            if string(matrix[x][y]) != "1" { continue }
            dp[x][y] = 1
            if x>0 && y>0 && dp[x][y]>0 {
                minSide := min(dp[x - 1][y], dp[x][y - 1])
                minSide = min(minSide, dp[x - 1][y - 1])
                dp[x][y] = minSide + 1
            }
            side = max(side, dp[x][y])
        }
    }
    
    return side * side
}

func min(a int, b int) int {
    if a > b {
        return b
    }
    return a
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
