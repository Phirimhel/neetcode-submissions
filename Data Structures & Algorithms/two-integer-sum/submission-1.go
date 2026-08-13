func twoSum(nums []int, target int) []int {
   needs := make(map[int]int)

   for i, n := range nums {
	  remainder := target - n 

	  if j, exists := needs[remainder]; exists {
		return []int{j, i}
	  }
	
	  needs[n] = i
   } 

return []int{}
}
