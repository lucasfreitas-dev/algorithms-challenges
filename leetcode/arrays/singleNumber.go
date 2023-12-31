// Single Number
// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/549/

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// You must implement a solution with a linear runtime complexity and use only constant extra space.

 

// Example 1:

// Input: nums = [2,2,1]
// Output: 1
// Example 2:

// Input: nums = [4,1,2,1,2]
// Output: 4
// Example 3:

// Input: nums = [1]
// Output: 1
 

// Constraints:

// 1 <= nums.length <= 3 * 104
// -3 * 104 <= nums[i] <= 3 * 104
// Each element in the array appears twice except for one element which appears only once.

func singleNumber(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
    // buildIn pdqsort: time O(n logn)
    sort.Ints(nums)
    
    // time O(n)
    for i := 0; i < len(nums)-1; i += 2 {
        if nums[i] != nums[i+1] {
            return nums[i]
        }
    }
    
    return nums[len(nums)-1]
}