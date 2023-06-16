// Intersection of Two Arrays II
// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/674/

// Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

 

// Example 1:

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]
// Example 2:

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]
// Explanation: [9,4] is also accepted.
 

// Constraints:

// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000
 
func intersect(nums1 []int, nums2 []int) []int {
    
    // space O(n)
    result := make([]int,0)
    
    // time O(nˆ2)
    for i := 0; i < len(nums1); i++ {
        for j := 0; j < len(nums2); j++ {
            if nums1[i] == nums2[j] {
                result = append(result, nums1[i])
                nums2 = append(nums2[:j],nums2[j+1:]...)
                break
            }
        }
    }
    
    return result
}