func twoSum(nums []int, target int) []int {
   needs := make(map[int]int)

   for i, n := range nums {
	  remainder := target - n 

	  if _, exists := needs[remainder]; exists {
		return []int{needs[remainder], i}
	  }
	
	  needs[n] = i
   } 

return []int{}
}
