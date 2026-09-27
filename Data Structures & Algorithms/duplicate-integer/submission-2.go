import "slices"

func hasDuplicate(nums []int) bool {
    slices.Sort(nums)

    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }
    return false
}
