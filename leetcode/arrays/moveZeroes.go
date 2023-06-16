// Move Zeroes
// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/567/

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

 

// Example 1:

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:

// Input: nums = [0]
// Output: [0]
 

// Constraints:

// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1

func moveZeroes(nums []int)  {
    
    countZeroes := 0
    
    // time O(n)
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[countZeroes] = nums[i]
            countZeroes++
        }   
    }
    
    // time O(n)
    for i := countZeroes; i < len(nums); i++ {
        nums[i] = 0
    }
}