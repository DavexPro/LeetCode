func twoSum(nums []int, target int) []int {
    result := make(map[int]int)
    for i:=0;i<len(nums);i++{
        rest :=  target - nums[i]
        if _, ok := result[rest]; ok == true {
			return []int {i, result[rest]}
		}
        result[nums[i]] = i
    }
    return []int{}
}
