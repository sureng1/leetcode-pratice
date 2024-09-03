/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 *
 * https://leetcode.cn/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (50.08%)
 * Likes:    957
 * Dislikes: 0
 * Total Accepted:    602.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '[5,2,3,1]'
 *
 * 给你一个整数数组 nums，请你将该数组升序排列。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,2,3,1]
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,1,1,2,0,0]
 * 输出：[0,0,1,1,2,5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 10^4
 * -5 * 10^4 <= nums[i] <= 5 * 10^4
 *
 *
 */

// @lc code=start
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	pa := sortArray(nums[0:mid])
	var pb []int
	if mid < len(nums) {
		pb = sortArray(nums[mid:])
	}
	var merged []int
	var i, j int
	for i < len(pa) && j < len(pb) {
		if pa[i] > pb[j] {
			merged = append(merged, pb[j])
			j++
		} else {
			merged = append(merged, pa[i])
			i++
		}
	}
	for ; i < len(pa); i++ {
		merged = append(merged, pa[i])
	}
	for ; j < len(pb); j++ {
		merged = append(merged, pb[j])
	}

	copy(nums, merged)
	// for i := 0; i < len(merged); i++ {
	// 	nums[i] = merged[i]
	// }
	return nums
}

// @lc code=end

