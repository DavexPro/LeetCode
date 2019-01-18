func countPrimes(n int) int {
    
    if n < 3 {
        return 0
    }
    
    check := make(map[int]bool, n)
    var prime []int
    
    for i := 2; i < n; i++ {
		if !check[i] {
			prime = append(prime, i)
		}else{
			continue
		}

		// 是当前素数的倍数的一定是合数
		for j := 2*i; j<n;j+=i {
			check[j] = true
		}
	}
    
    return len(prime)
}
